package controller

import (
	"fmt"
	"net/http"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserControllerContainer interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

type UserControllerDependency struct {
	service services.UserServiceContainer
}

func NewUserControllerDependency(service services.UserServiceContainer) *UserControllerDependency {
	return &UserControllerDependency{
		service: service,
	}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *UserControllerDependency) Signup(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
	}

	if err2 := c.service.CreateUser(&user); err2 != nil {
		ctx.Error(err2)
	}

	// Create web token for user.
	email := user.Email
	token, err3 := helpers.GenerateToken(email)
	if err3 != nil {
		ctx.Error(err3)
	}

	ctx.JSON(http.StatusCreated, gin.H{"msg": "User created", "token": token})
	// ctx.Set("token", token)
	fmt.Print(token)
}

func (c *UserControllerDependency) Login(ctx *gin.Context) {
	var user UserLogin
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
	}

	email := user.Email
	pass := user.Password

	dbUser, err := c.service.GetUserByEmail(email)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "User not found"})
	}

	hashed := dbUser.Password

	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := helpers.GenerateToken(email)
	if err != nil {
		ctx.Error(err)
	}
	ctx.Set("token", token)
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid password"})
}
