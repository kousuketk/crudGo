package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/repository"
)

type AuthController struct{}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password"`
}

var userRepo repository.UserRepository

func (a *AuthController) Login(c *gin.Context) {
	var params LoginParam
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := userRepo.GetUserByEmail(params.Email)
	if user.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	flag, err := middlewares.VerifyPassword(user, params.Password)
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is not correct"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := middlewares.CreateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":  token,
		"user": user,
	})
}

func (a *AuthController) Logout(c *gin.Context) {
	// jwtをinvalidにする
	c.JSON(http.StatusOK, gin.H{"auth": "Logout"})
}
