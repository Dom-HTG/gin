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
	HomeHandler()
}

type ControllerDependencies struct {
	service services.ServiceContainer
}

func NewControllerDependencies(service services.ServiceContainer) *ControllerDependencies {
	return &ControllerDependencies{
		service: service,
	}
}

func (cd *ControllerDependencies) HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "This is the home page."})
}

func (cd *ControllerDependencies) ListProducts(ctx *gin.Context) {
	products, err := cd.service.GetAllProducts()
	if err != nil {
		ctx.Error(err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (cd *ControllerDependencies) ListProduct(ctx *gin.Context) {
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

func (cd *ControllerDependencies) AddProduct(ctx *gin.Context) {
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

func (cd *ControllerDependencies) UpdateProduct(ctx *gin.Context) {
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

func (cd *ControllerDependencies) DeleteProduct(ctx *gin.Context) {
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
