package controller

import (
	"net/http"
	"strconv"

	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/services"
	"github.com/gin-gonic/gin"
)

type ProductContainer interface {
	ListProducts()
	ListProduct(id string)
	AddProduct(product *models.Product)
	UpdateProduct(id string, product *models.Product)
	DeleteProduct(id string)
}

type ProductControllerDependencies struct {
	service services.ProductServiceContainer
}

func NewControllerDependencies(service services.ProductServiceContainer) *ProductControllerDependencies {
	return &ProductControllerDependencies{
		service: service,
	}
}

func HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "This is the home page."})
}

func (cd *ProductControllerDependencies) ListProducts(ctx *gin.Context) {
	products, err := cd.service.GetAllProducts()
	if err != nil {
		ctx.Error(err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (cd *ProductControllerDependencies) ListProduct(ctx *gin.Context) {
	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.Error(err)
	}

	product, err := cd.service.GetProductByID(id)
	if err != nil {
		ctx.Error(err)
	}
	ctx.JSON(http.StatusOK, product)
}

func (cd *ProductControllerDependencies) AddProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.Error(err)
	}

	errr := cd.service.AddProduct(product)
	if err != nil {
		ctx.Error(errr)
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "success", "log": "new resource created"})
}

func (cd *ProductControllerDependencies) UpdateProduct(ctx *gin.Context) {
	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.Error(err)
	}

	var product models.Product
	errr := ctx.BindJSON(&product)
	if err != nil {
		ctx.Error(errr)
	}

	errrr := cd.service.UpdatedProduct(id, product)
	if errrr != nil {
		ctx.Error(errrr)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "log": "resource updated successfully"})
}

func (cd *ProductControllerDependencies) DeleteProduct(ctx *gin.Context) {
	idstring := ctx.Param("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		ctx.Error(err)
	}

	errr := cd.service.DeleteProduct(id)
	if errr != nil {
		ctx.Error(errr)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Success", "log": "resource deleted successfully"})
}
