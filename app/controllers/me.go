package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/services"
)

type MeController struct{}

type MeParam struct {
	Name             string `json:"name" binding:"required,min=1,max=50"`
	SelfIntroduction string `json:"selfIntroduction"`
	Email            string `json:"email" binding:"required"`
	PasswordDigest   string `json:"password"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phoneNumber"`
}

func (m *MeController) Index(c *gin.Context) {
	// cをもらってuserを返すmiddlewareを作成する
	userId := middlewares.Authorization(c)
	userServices := services.UserServices{}
	user, err := userServices.GetUserById(userId)
	if user.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (m *MeController) UpdateMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"me": "me UpdateMe"})
}

func (m *MeController) DeleteMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"me": "me DeleteMe"})
}
