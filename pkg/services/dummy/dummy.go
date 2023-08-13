package dummy

import (
	"github.com/dgalifi/go-webapp/pkg/config"
)

type DummyService interface {
	DoSomething() string
}

type dummyService struct {
	config config.Config
}

func NewDummyService(cfg config.Config) DummyService {
	return &dummyService{config: cfg}
}

func (ds *dummyService) DoSomething() string {
	return ds.config.GreetingsMessage
}
