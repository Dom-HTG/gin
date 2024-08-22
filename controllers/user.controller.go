package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dom-HTG/gin/helpers"
	"github.com/Dom-HTG/gin/models"
	"github.com/Dom-HTG/gin/services"
	"github.com/Dom-HTG/gin/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	//create session for user
	sessionID := uuid.NewString()
	sess := sessions.Default(ctx)
	sess.Set("session_id", sessionID)
	sess.Set("email", user.Email)

	if err := sess.Save(); err != nil {
		log.Fatal(err)
	}
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

	//create session for user
	sessionID := uuid.NewString()

	sessionToken, err := utils.GenerateJWTSession(sessionID, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.Default(ctx)
	session.Set("session_token", sessionToken)

	if err := session.Save(); err != nil {
		log.Fatal(err)
	}
}
