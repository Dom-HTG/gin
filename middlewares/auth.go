package middlewares

import (
	"net/http"
	"strings"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMW checks and verifies the token in the Authorization header.
func JWTAuthMW() gin.HandlerFunc {
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

		//Map Claim & Validate Token.
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.JSON(http.StatusOK, claims)
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	}

}
