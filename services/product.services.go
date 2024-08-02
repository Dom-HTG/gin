package services

import (
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/repository"
)

type ProductServiceContainer interface {
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(product models.Product) error
	UpdatedProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type ProductServiceDependency struct {
	repo repository.ProductRepositoryContainer
}

func NewServiceDependency(repo repository.ProductRepositoryContainer) *ProductServiceDependency {
	return &ProductServiceDependency{
		repo: repo,
	}
}

func (s *ProductServiceDependency) GetProductByID(id int) (models.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *ProductServiceDependency) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ProductServiceDependency) AddProduct(product models.Product) error {
	err := s.repo.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceDependency) UpdatedProduct(id int, product models.Product) error {
	err := s.repo.UpdateProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceDependency) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
