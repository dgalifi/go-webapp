package endpoints

import (
	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/dgalifi/go-webapp/pkg/services/dummy"
	"github.com/gin-gonic/gin"
)

type EndpointProvider struct {
	Config       config.Config
	DummyService dummy.DummyService
}

func (ep EndpointProvider) DoSomethingDummy(c *gin.Context) {
	ep.DummyService.DoSomething(c)
}
