package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	sessionID string
	jwt.StandardClaims
}

type claimsConstant struct {
	expiresAt     int64
	sessionSecret []byte
}

var Constant = &claimsConstant{
	expiresAt:     time.Now().Add(24 * time.Hour).Unix(),
	sessionSecret: []byte(os.Getenv("SESSION_SECRET")),
}

func generateJWTSession(sessionID string) (string, error) {

	//build new claims for JWT token.
	sessionClaims := &Claims{
		sessionID: sessionID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Constant.expiresAt,
		},
	}

	//build new token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims)

	//sign new token.
	signedToken, err := token.SignedString(Constant.sessionSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func verifyJWTSession(token string) (string, error) {
	tokenClaims := &Claims{}

	parsed_token, err := jwt.ParseWithClaims(token, tokenClaims, func(tk *jwt.Token) (interface{}, error) {
		return Constant.sessionSecret, nil
	})

	if err != nil {
		return "", errors.New("failed to parse token")
	}

	if parsed_token.Valid {
		return tokenClaims.sessionID, nil
	}
	return "", errors.New("token is not valid")
}
