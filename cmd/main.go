package main

import (
	"fmt"
	"log"
	"os"

	controller "github.com/Dom-HTG/gin/controllers"
	"github.com/Dom-HTG/gin/middlewares"
	"github.com/Dom-HTG/gin/repository"
	"github.com/Dom-HTG/gin/services"
	"github.com/Dom-HTG/gin/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	//Load environment variables.
	err := godotenv.Load("main.env", "sessions.env")
	if err != nil {
		log.Fatalf("error loading environment variables %v", err)
	}

	//database connection.
	db, err := utils.InitializeDatabase()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}

	//init redis store.
	redisStore, err := utils.InitRedisStore("localhost:8080")
	if err != nil {
		log.Fatalf("error initializing redis store: %v", err)
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
	access.Use(sessions.Sessions("login-session", redisStore))
	{
		access.POST("/signup", userController.Signup)
		access.POST("/login", middlewares.Authenticate(), userController.Login)
	}

	router.Run(port)
	fmt.Printf("server is runing on port %s \n", port)

}
