package services

import (
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/repository"
)

type ServiceContainer interface {
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(product models.Product) error
	UpdatedProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type ServiceDependency struct {
	repo repository.RepositoryContainer
}

func NewServiceDependency(repo repository.RepositoryContainer) *ServiceDependency {
	return &ServiceDependency{
		repo: repo,
	}
}

func (s *ServiceDependency) GetProductByID(id int) (models.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *ServiceDependency) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ServiceDependency) AddProduct(product models.Product) error {
	err := s.repo.AddProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceDependency) UpdatedProduct(id int, product models.Product) error {
	err := s.repo.UpdatedProduct(id, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceDependency) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
