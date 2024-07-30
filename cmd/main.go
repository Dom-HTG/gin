package main

import (
	"log"

	controller "github.com/Dom-HTG/gin/controllers"
	"github.com/Dom-HTG/gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//Instance for injecting dependencies
	productSample := &controller.ProductSample{}

	//Instantiate gin Router and group routes.
	router := gin.Default()

	router.GET("/home", productSample.HomeHandler)
	router.GET("/products", productSample.ListProducts)
	router.GET("/products/:id", productSample.ListProduct)

	//All endpoints that allows requests to be mutated are 'protected'.
	protected := router.Group("/api/protected")
	{
		protected.POST("/products", productSample.AddProduct)
		protected.PUT("/products/:id", productSample.UpdateProduct)
		protected.DELETE("/products/:id", productSample.DeleteProduct)
	}

	router.Run(models.Config.Port)

}
