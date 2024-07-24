package main

import (
	"log"
	"os"

	"github.com/Dom-HTG/gin/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//get environment variables.
	PORT := os.Getenv("PORT")

	//Instantiate Router
	router := gin.Default()

	router.GET("/home", handlers.HomeHandler)
	router.GET("/products", handlers.ListProducts)
	router.GET("/products/:id", handlers.ListProduct)
	router.POST("/products", handlers.AddProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	router.Run(PORT)
}
