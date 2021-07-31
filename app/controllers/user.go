package controllers

import (
	"app/middlewares"
	"app/models"
	"net/http"

	"github.com/git-gonic/gin"
)

type UserController struct{}

type UserParam struct {
	Name  string `json:"name" binding:"required,min=1,max=50"`
	Email string `json:"email" binding:"required`
}

// ユーザー一覧
func (self *UserController) Index(c *gin.Context) {
	users, err := models.GetUsers(middlewares.DB())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user search failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
