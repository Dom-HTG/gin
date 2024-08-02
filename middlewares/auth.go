package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate middleware checks and verifies the token in the Authorization header.
func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Check if Token exists.
		BearerToken := ctx.GetHeader("Authorization")
		if BearerToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header."})
			ctx.Abort()
			return
		}
		tokenstring := strings.TrimPrefix(BearerToken, "Bearer")

		//Parse token.
		token, err := jwt.Parse(tokenstring, helpers.VerifyToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot Parse token string into JWT Object."})
			ctx.Abort()
		}

		//Map Claims & Validate Token.
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.JSON(http.StatusOK, gin.H{"msg": "token is valid", "claims": claims})
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}
}

// structure to hold token claims.
type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateToken(email string) (string, error) {
	exp := time.Now().Add(24 * time.Hour)

	newClaims := &claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	signed_token, err := token.SignedString(helpers.JWTSECRET)
	if err != nil {
		return "", err
	}
	return signed_token, nil
}
