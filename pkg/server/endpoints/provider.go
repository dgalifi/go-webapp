package endpoints

import (
	"net/http"

	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/dgalifi/go-webapp/pkg/services/dummy"
	"github.com/gin-gonic/gin"
)

type EndpointProvider struct {
	Config       config.Config
	DummyService dummy.DummyService
}

func (ep EndpointProvider) HealthCheck(c *gin.Context) {
	statusCode := http.StatusOK

	retMsg := "Healthy"

	c.JSON(statusCode, gin.H{
		"message": retMsg,
	})
}

func (ep EndpointProvider) DoSomethingDummy(c *gin.Context) {
	statusCode := http.StatusOK

	retMsg := ep.DummyService.DoSomething()

	c.JSON(statusCode, gin.H{
		"message": retMsg,
	})
}
