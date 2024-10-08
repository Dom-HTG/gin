package middlewares

import (
	"net/http"
	"strings"

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
		tokenstring := strings.TrimPrefix(BearerToken, "Bearer ")

		//Parse token.
		token, err := jwt.Parse(tokenstring, helpers.VerifyToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Cannot Parse token string into JWT Object."})
			ctx.Abort()
		}

		//Map Claims & Validate Token.
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			ctx.Set("claims", claims)
			ctx.Next()
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
	}
}
