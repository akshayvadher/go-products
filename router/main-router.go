package router

import (
	"github.com/gin-gonic/gin"
	"training/go-products/products"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	productsRouterGroup := router.Group("/products")
	{
		mappings(productsRouterGroup)
	}
	return router

}

func mappings(productsRouterGroup *gin.RouterGroup) {
	productsRouterGroup.GET("/", products.GetStatus)
	productsRouterGroup.POST("/", products.SaveProduct)
	productsRouterGroup.GET("/:id", products.GetById)
	productsRouterGroup.GET("/list", products.GetAllProductsController)
}
