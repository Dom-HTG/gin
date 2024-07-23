package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "This is the home page."})
}

func ListProducts(ctx *gin.Context) {

}
