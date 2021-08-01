package controllers

import (
	"net/http"
	"strconv"

	"github.com/kousuketk/crudGo/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserParam struct {
	Name  string `json:"name" binding:"required,min=1,max=50"`
	Email string `json:"email" binding:"required"`
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
