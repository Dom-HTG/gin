package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/Dom-HTG/gin/models"
	"github.com/dgrijalva/jwt-go"
)

var JWTSECRET = os.Getenv("JWTSECRET")

// helper function to check token validity.
func VerifyToken(token *jwt.Token) (interface{}, error) {
	_, verified := token.Method.(*jwt.SigningMethodHMAC)
	if verified {
		return models.Config.JWTSECRET, nil
	}
	return nil, fmt.Errorf("unexpected algo: %v", token.Header["alg"])
}

// structure to hold token claims.
type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	newClaims := &claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	signed_token, err := token.SignedString(JWTSECRET)
	if err != nil {
		return "", err
	}
	return signed_token, nil
}
