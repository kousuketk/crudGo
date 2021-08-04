package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/services"
)

type AuthController struct{}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"Password"`
}

func (self *AuthController) Login(c *gin.Context) {
	// もしjwtがあったら, それで検証する
	// →userが見つかる→alredy login, userが見つからない→以下の処理へ以降
	// param.email, passを見て検証する, useridを保持する
	// findbyemai, passを見て検証する
	var param LoginParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userServices := services.UserServices{}
	user, err := userServices.GetUserByEmail(param.Email)
	if user.IsEmpty() == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	flag, _ := middlewares.VerifyPassword(user, param.Password)
	if flag == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is not correct"})
		return
	}

	token, err3 := middlewares.CreateToken(strconv.FormatUint(uint64(user.ID), 10))
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt":  token,
		"user": user,
	})
}

func (self *AuthController) Logout(c *gin.Context) {
	// jwtをinvalidにする
	c.JSON(http.StatusOK, gin.H{"auth": "Logout"})
}
