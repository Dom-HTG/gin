package utils

import (
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
	conn.AutoMigrate(&models.Product{})
	return conn, nil
}
