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
	users := userServices.GetUsers()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (self *UserController) GetUser(c *gin.Context) {
	userServices := services.UserServices{}
	ID := c.Params.ByName("id")
	userID, _ := strconv.Atoi(ID)
	user := userServices.GetUserById(userID)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
