package main

import (
	"log"

	controller "github.com/Dom-HTG/gin/controllers"
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/repository"
	"github.com/Dom-HTG/gin/services"
	"github.com/Dom-HTG/gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//database connection.
	db, err := utils.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	//Instance for injecting dependencies
	productRepo := repository.NewRepoDependencies(db)
	productService := services.NewServiceDependency(productRepo)
	productController := controller.NewControllerDependencies(productService)

	//Instantiate gin Router and group routes.
	router := gin.Default()

	router.GET("/home", productController.HomeHandler)
	router.GET("/products", productController.ListProducts)
	router.GET("/products/:id", productController.ListProduct)

	//All endpoints that allows requests to be mutated are 'protected'.
	protected := router.Group("/api/protected")
	{
		protected.POST("/products", productController.AddProduct)
		protected.PUT("/products/:id", productController.UpdateProduct)
		protected.DELETE("/products/:id", productController.DeleteProduct)
	}

	router.Run(models.Config.Port)

}
