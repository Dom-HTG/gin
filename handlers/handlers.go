package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/Dom-HTG/gin/models"
	"github.com/gin-gonic/gin"
)

func HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "This is the home page."})
}

func ListProducts(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(sample)
	ctx.IndentedJSON(http.StatusOK, sample)
}

func ListProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		log.Fatal(err)
	}

	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	for _, product := range sample {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
			return
		}
	}
}

func AddProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		log.Fatal(err)
	}

	var newProduct models.Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sample = append(sample, newProduct)
	ctx.JSON(http.StatusCreated, sample)
}

func UpdateProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		log.Fatal(err)
	}

	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	var UpdatedProduct models.Product
	if err := ctx.BindJSON(&UpdatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, product := range sample {
		if product.ID == id {
			product = UpdatedProduct
		}
		return
	}
	ctx.JSON(http.StatusOK, sample)
}

func DeleteProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		log.Fatal(err)
	}

	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	for i, product := range sample {
		if product.ID == id {
			sample = append(sample[:i], sample[i+1:]...)
		}
		return
	}
	ctx.JSON(http.StatusOK, sample)
}
