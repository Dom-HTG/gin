package models

import "os"

type Product struct {
	ID          int                `json:"id" gorm:"primary key"`
	Title       string             `json:"title" gorm:"column:product_title"`
	Description string             `json:"description" gorm:"column:product_description"`
	Price       float64            `json:"price" gorm:"column:product_price"`
	Tags        []string           `json:"tags" gorm:"column:product_tags;type:text[]"`
	Brand       string             `json:"brand" gorm:"column:product_brand"`
	Dimensions  map[string]float64 `json:"dimensions" gorm:"column:product_dimensions;type:jsonb"`
}

type User struct {
	FirstName string `json:"firstname" gorm:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname"`
	Email     string `json:"email" gorm:"email"`
	Password  string `json:"password" gorm:"-"`
}

type ProductStore struct {
	Products []Product `json:"products"`
}

var Tags = []string{}
var Dimensions = map[string]float64{}

var Config = struct {
	Port      string
	JWTSECRET string
}{
	Port:      os.Getenv("PORT"),
	JWTSECRET: os.Getenv("JWTSECRET"),
}
