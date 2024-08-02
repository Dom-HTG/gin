package repository

import (
	"github.com/Dom-HTG/gin/models"
	"gorm.io/gorm"
)

type ProductRepositoryContainer interface {
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(product models.Product) error
	UpdateProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type RepoDependencies struct {
	DBConn *gorm.DB
}

func NewRepoDependencies(DBConn *gorm.DB) *RepoDependencies {
	return &RepoDependencies{
		DBConn: DBConn,
	}
}

func (r *RepoDependencies) GetProductByID(id int) (models.Product, error) {
	var product models.Product

	results := r.DBConn.Where("id = ?", id).Find(&product)
	if results.Error != nil {
		return models.Product{}, results.Error
	}
	return product, nil
}

func (r *RepoDependencies) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	results := r.DBConn.Find(&products)
	if results.Error != nil {
		return []models.Product{}, results.Error
	}
	return products, nil
}

func (r *RepoDependencies) AddProduct(product models.Product) error {
	results := r.DBConn.Create(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

func (r *RepoDependencies) UpdateProduct(id int, product models.Product) error {
	results := r.DBConn.Where("id = ?", id).Updates(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

func (r *RepoDependencies) DeleteProduct(id int) error {
	var product models.Product

	results := r.DBConn.Where("id = ?", id).Delete(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
