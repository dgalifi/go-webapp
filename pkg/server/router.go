package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/dgalifi/go-webapp/pkg/server/endpoints"
	"github.com/dgalifi/go-webapp/pkg/services/dummy"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	GetEngine() *gin.Engine
}

type server struct {
	config config.Config
	engine *gin.Engine
}

func NewServer(config config.Config, ds dummy.DummyService) Server {
	engine := gin.Default()
	addRoutes(engine, config, ds)

	return &server{engine: engine, config: config}
}

func addRoutes(
	engine *gin.Engine,
	config config.Config,
	ds dummy.DummyService) {

	ep := endpoints.EndpointProvider{
		Config:       config,
		DummyService: ds,
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/dummy", ep.DoSomethingDummy)

	engine.GET("/health", ep.HealthCheck)
}

// Start will bind a Gin router to a TCP port and start
// the web server
func (server *server) Start() error {
	if server.config.WebServerPort == "" {
		return errors.New("env PORT is not defined, it must be a valid TCP port to start the server")
	}

	s := &http.Server{
		Addr:         ":" + server.config.WebServerPort,
		Handler:      server.engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// ListenAndServe will block until it is stopped, then it returns an error.
	return s.ListenAndServe()
}

func (server *server) GetEngine() *gin.Engine {
	return server.engine
}
