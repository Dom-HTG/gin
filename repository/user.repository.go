package repository

import "github.com/Dom-HTG/gin/models"

type RepositoryContainer interface {
	GetProductByID(id int) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	AddProduct(product models.Product) error
	UpdatedProduct(id int, product models.Product) error
	DeleteProduct(id int) error
}

type UserRepository struct {
	DB *db.DB
}

func (r *UserRepository) GetProductByID(id int) (models.Product, error) {
	//database logic here.
}

func (r *UserRepository) GetAllProducts() ([]models.Product, error) {
	//database logic here.
}

func (r *UserRepository) AddProduct(product models.Product) error {
	//database logic here.
}

func (r *UserRepository) UpdateProduct(id int, product models.Product) error {
	//database logic here.
}

func (r *UserRepository) DeleteProduct(id int) error {
	//database logic here.
}
