package main

import (
	"github.com/Dom-HTG/gin/handlers"
	"github.com/gin-gonic/gin"
)

type ProductStore interface {
	Get()
	List()
	Add()
	Update()
	Delete()
}

func main() {
	router := gin.Default()

	router.GET("/home", handlers.HomeHandler)
	router.GET("/products", handlers.ListProducts)
	router.GET("/products/:id", handlers.ListProduct)
	router.POST("/products/:id", handlers.AddProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	router.Run(":4030")
}
