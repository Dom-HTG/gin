package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

// helper function to fetch dummydata from third party api.
func DummyData() ([]models.Product, error) {
	URL := "https://dummyjson.com/products"
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newProduct []models.Product
	if err := json.Unmarshal(response, &newProduct); err != nil {
		return nil, err
	}

	return newProduct, nil
}
