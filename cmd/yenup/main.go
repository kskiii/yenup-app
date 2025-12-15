package main

import (
	"fmt"
	"log"
	"yenup/internal/config"
	"yenup/internal/registory"

	"github.com/gin-gonic/gin"
)

func main() {

	// load config from config.go
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// create registory from registory.go
	reg, err := registory.NewRegistory(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// create app handler from registory
	appHandler := reg.AppHandler

	// create router from gin
	r := gin.Default()

	// register routes from route.go
	appHandler.RegisterRoutes(r)

	// run server
	r.Run(fmt.Sprintf(":%s", cfg.AppPort))

}
