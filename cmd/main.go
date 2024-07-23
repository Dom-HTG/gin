package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Dom-HTG/gin/models"
	"github.com/gin-gonic/gin"
)

type ProductStore interface {
	Get(id int)
	List()
	Add()
	Update()
	Delete()
}

func dummydata() ([]models.Product, error) {
	URL := "https://dummyjson.com/products"
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newProduct []models.Product
	if err := json.Unmarshal(response, &newProduct); err != nil {
		return nil, err
	}

	return newProduct, nil
}

func main() {
	router := gin.Default()

	router.GET("/home", handlers.homeHandler)
	router.GET("/products", handlers.ListProducts)
	router.GET("/products/:id", handlers.ListProduct)
	router.POST("/products/:id", handlers.AddProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)

	router.Run(":4030")
}
