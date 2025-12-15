package main

import (
	"fmt"
	"log"
	"yenup/internal/config"
	"yenup/internal/registry"

	"github.com/gin-gonic/gin"
)

func main() {

	// load config from config.go
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// create registry from registry.go
	reg, err := registry.NewRegistry(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// create app handler from registry
	appHandler := reg.AppHandler

	// create router from gin
	r := gin.Default()

	// register routes from route.go
	appHandler.RegisterRoutes(r)

	// run server
	r.Run(fmt.Sprintf(":%s", cfg.AppPort))

}
