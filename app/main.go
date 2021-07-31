package main

import (
	"net/http"

	"app/middlewares"

	"app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// gormのDB接続
	middlewares.Setup()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r := router.Group("/api/v1")
	controllers.Setup(r)

	router.Run()
}
