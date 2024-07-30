package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Dom-HTG/gin/models"
	"github.com/dgrijalva/jwt-go"
)

var JWTSECRET = os.Getenv("JWTSECRET")

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

func VerifyToken(token *jwt.Token) (interface{}, error) {
	_, verified := token.Method.(*jwt.SigningMethodHMAC)
	if verified {
		return models.Config.JWTSECRET, nil
	}
	return nil, fmt.Errorf("unexpected algo: %v", token.Header["alg"])
}
