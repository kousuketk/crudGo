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
	PasswordDigest   string `json:"passwordDigest"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phoneNumber"`
}

// ユーザー一覧
func (self *UserController) Index(c *gin.Context) {
	userServices := services.UserServices{}
	users, err := userServices.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (self *UserController) GetUser(c *gin.Context) {
	ID := c.Params.ByName("id")
	userID, _ := strconv.Atoi(ID)
	userServices := services.UserServices{}
	user, err := userServices.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (self *UserController) CreateUser(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "user create failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
