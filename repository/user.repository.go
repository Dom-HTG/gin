package repository

import (
	"github.com/Dom-HTG/gin/models"
	"gorm.io/gorm"
)

type UserRepositoryContainer interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

type UserRepoDependency struct {
	DBConn *gorm.DB
}

func (d *UserRepoDependency) CreateUser(user *models.User) error {
	tx := d.DBConn.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d *UserRepoDependency) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	tx := d.DBConn.Where("email = ?", email).First(&user)
	if tx.Error == nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *UserRepoDependency) UpdateUser(id int, user *models.User) error {
	tx := d.DBConn.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d *UserRepoDependency) DeleteUser(id int) error {
	tx := d.DBConn.Where("id = ?", id).Delete(&models.User{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
