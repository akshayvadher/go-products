package main

import (
	"training/go-products/persistance"
	"training/go-products/router"
	"training/go-products/validation"
)

func main() {
	validation.InitValidator()

	persistance.CreateConnection()

	routes := router.InitRoutes()

	routes.Run("localhost:8080")
}
