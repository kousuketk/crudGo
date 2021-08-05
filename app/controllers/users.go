package controllers

import (
	"net/http"
	"strconv"

	"github.com/kousuketk/crudGo/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserParam struct {
	Name             string `json:"name" binding:"required,min=1,max=50"`
	SelfIntroduction string `json:"selfIntroduction"`
	Email            string `json:"email" binding:"required"`
	PasswordDigest   string `json:"password"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phoneNumber"`
}

// ユーザー一覧
func (u *UserController) Index(c *gin.Context) {
	userServices := services.UserServices{}
	users, err := userServices.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	userId, _ := strconv.Atoi(id)
	userServices := services.UserServices{}
	user, err := userServices.GetUserById(userId)
	if user.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) CreateUser(c *gin.Context) {
	var param UserParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userServices := services.UserServices{}
	user, err := userServices.CreateUser(
		param.Name,
		param.SelfIntroduction,
		param.Email,
		param.PasswordDigest,
		param.Address,
		param.PhoneNumber)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
