package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kousuketk/crudGo/app/controllers"
	"github.com/kousuketk/crudGo/app/middlewares"
)

func Router() *gin.Engine {
	// gormのDB接続
	middlewares.Setup()
	router := gin.Default()

	// 環境変数の設定
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println(err)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r := router.Group("/api/v1")
	controllers.Setup(r)
	return router
}

func main() {
	Router().Run()
}
