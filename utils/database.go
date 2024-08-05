package utils

import (
	"fmt"
	"os"

	"github.com/Dom-HTG/gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initialize and return database connection.
func InitializeDatabase() (*gorm.DB, error) {
	PG_URL := os.Getenv("PG_URL")
	conn, err := gorm.Open(postgres.Open(PG_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Print("POSTGRES database connection established")

	conn.AutoMigrate(&models.Product{}, &models.User{})

	return conn, nil
}
