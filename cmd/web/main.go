package main

import (
	"fmt"
	"log"

	"github.com/dgalifi/go-webapp/pkg/config"
	"github.com/dgalifi/go-webapp/pkg/server"
	"github.com/dgalifi/go-webapp/pkg/services/dummy"
)

func main() {
	cfg := config.Config{
		WebServerPort:    "8080",
		GreetingsMessage: "Hello World",
	}

	ds := dummy.NewDummyService(cfg)

	s := server.NewServer(cfg, ds)

	fmt.Printf("### Webserver running on port %v ###\n", cfg.WebServerPort)
	log.Fatal(s.Start())
}
