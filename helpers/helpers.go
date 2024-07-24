package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Dom-HTG/gin/models"
)

// helper function to fetch dummydata from third party api.
func DummyData() ([]models.Product, error) {
	res, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var sampleProduct models.ProductStore

	if err := json.Unmarshal(response, &sampleProduct); err != nil {
		return nil, err
	}
	return sampleProduct.Products, nil
}
