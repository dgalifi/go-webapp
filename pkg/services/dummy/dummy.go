package dummy

import (
	"net/http"

	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/gin-gonic/gin"
)

type DummyService interface {
	DoSomething(ginContext *gin.Context)
}

type dummyService struct {
	config config.Config
}

func NewDummyService(cfg config.Config) DummyService {
	return &dummyService{config: cfg}
}

func (ds *dummyService) DoSomething(c *gin.Context) {

	statusCode := http.StatusOK

	c.JSON(statusCode, gin.H{
		"message": ds.config.GreetingsMessage,
	})
}
