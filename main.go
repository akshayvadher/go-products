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

	err := DbMigration()
	if err != nil {
		return
	}

	routes := router.InitRoutes()

	startServer(routes)
}

func startServer(routes *gin.Engine) {
	var port = strconv.Itoa(config.AppConf.Server.Port)
	log.Println("Starting the server on port: " + port)
	err := routes.Run("127.0.0.1:" + port)
	if err != nil {
		log.Println("Error starting server: " + err.Error())
		return
	}
}
