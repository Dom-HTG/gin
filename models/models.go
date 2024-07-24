package models

import "os"

type Product struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Tags        []string           `json:"tags"`
	Brand       string             `json:"brand"`
	Dimensions  map[string]float64 `json:"dimensions"`
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
