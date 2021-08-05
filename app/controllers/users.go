package controllers

import (
	"net/http"
	"strconv"

	"github.com/kousuketk/crudGo/app/models"
	"github.com/kousuketk/crudGo/app/serializers"
	"github.com/kousuketk/crudGo/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Index(c *gin.Context) {
	userServices := services.UserServices{}
	users, err := userServices.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	result := serializers.UserSliceSerialize(users)

	c.JSON(http.StatusOK, gin.H{"users": result})
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
	result := serializers.UserSerialize(user)

	c.JSON(http.StatusOK, gin.H{"user": result})
}

func (u *UserController) CreateUser(c *gin.Context) {
	var userParams models.User
	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userServices := services.UserServices{}
	user, err := userServices.CreateUser(userParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
