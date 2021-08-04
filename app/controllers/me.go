package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (self *MeController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"me": "me index"})
}

func (self *MeController) UpdateMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"me": "me UpdateMe"})
}

func (self *MeController) DeleteMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"me": "me DeleteMe"})
}
