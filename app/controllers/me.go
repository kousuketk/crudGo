package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kousuketk/crudGo/app/middlewares"
	"github.com/kousuketk/crudGo/app/models"
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

var meServices services.MeServices

func (m *MeController) Index(c *gin.Context) {
	userId, flag := middlewares.Authorization(c)
	if !flag {
		return
	}
	user, err := meServices.Index(userId)
	if user.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found."})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (m *MeController) UpdateMe(c *gin.Context) {
	userId, flag := middlewares.Authorization(c)
	if !flag {
		return
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := meServices.UpdateMe(userId, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success updated."})
}

func (m *MeController) DeleteMe(c *gin.Context) {
	userId, flag := middlewares.Authorization(c)
	if !flag {
		return
	}
	err := meServices.DeleteMe(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success deleted."})
}
