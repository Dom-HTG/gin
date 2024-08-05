package main

import (
	"fmt"
	"log"

	controller "github.com/Dom-HTG/gin/controllers"
	"github.com/Dom-HTG/gin/middlewares"
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

	//product layer dependencies.
	productRepo := repository.NewRepoDependencies(db)
	productService := services.NewServiceDependency(productRepo)
	productController := controller.NewControllerDependencies(productService)

	//user Layerdependencies.
	userRepo := repository.NewUserRepoDependency(db)
	userService := services.NewUserServiceDependency(userRepo)
	userController := controller.NewUserControllerDependency(userService)

	//Instantiate gin Router and group routes.
	router := gin.Default()

	router.GET("/home", controller.HomeHandler)

	router.GET("/products", productController.ListProducts)
	router.GET("/products/:id", productController.ListProduct)

	//All endpoints that allows mutation of data are 'protected'.
	protected := router.Group("/api/protected")
	{
		protected.POST("/products", middlewares.Authenticate(), productController.AddProduct)
		protected.PUT("/products/:id", middlewares.Authenticate(), productController.UpdateProduct)
		protected.DELETE("/products/:id", middlewares.Authenticate(), productController.DeleteProduct)
	}

	access := router.Group("/api/access")
	{
		access.POST("/signup", userController.Signup)
		access.POST("/login", middlewares.Authenticate(), userController.Login)
	}

	router.Run(models.Config.Port)
	fmt.Printf("server started on port %s\n", models.Config.Port)

}
