package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"training/go-products/config"
	"training/go-products/persistance"
	"training/go-products/router"
	"training/go-products/validation"
)

func main() {
	config.ReadConfig()

	validation.InitValidator()

	persistance.CreateConnection()

	routes := router.InitRoutes()

	startServer(routes)
}

func startServer(routes *gin.Engine) {
	var port = strconv.Itoa(config.AppConf.Server.Port)
	log.Println("Starting the server on port" + port)
	err := routes.Run(":" + port)
	if err != nil {
		return
	}
}
