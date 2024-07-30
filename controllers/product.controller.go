package controller

import (
	"net/http"
	"strconv"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/Dom-HTG/gin/models"
	"github.com/gin-gonic/gin"
)

type ProductContainer interface {
	ListProducts() ([]models.Product, error)
	ListProduct(id string) ([]models.Product, error)
	AddProduct(product *models.Product) error
	UpdateProduct(id string, product *models.Product) error
	DeleteProduct(id string) error
	HomeHandler() any
}

type ProductSample struct{}

func (ps *ProductSample) HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "This is the home page."})
}

func (ps *ProductSample) ListProducts(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, sample)
}

func (ps *ProductSample) ListProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	for _, product := range sample {
		if product.ID == id {
			ctx.JSON(http.StatusOK, product)
		}
	}
}

func (ps *ProductSample) AddProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var newProduct models.Product
	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	sample = append(sample, newProduct)
	ctx.JSON(http.StatusCreated, sample)
}

func (ps *ProductSample) UpdateProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
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
	}
	ctx.JSON(http.StatusOK, sample)
}

func (ps *ProductSample) DeleteProduct(ctx *gin.Context) {
	sample, err := helpers.DummyData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
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
	}
	ctx.JSON(http.StatusOK, sample)
}
