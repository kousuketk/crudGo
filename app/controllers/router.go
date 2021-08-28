package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kousuketk/crudGo/app/middlewares"
)

func Setup(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		u := UserController{}
		users.GET("", u.Index)
		users.GET("/:id", u.GetUser)
		users.POST("", u.CreateUser)
	}
	auth := r.Group("")
	{
		a := AuthController{}
		auth.POST("/login", a.Login)
		auth.POST("/logout", a.Logout)
	}
	me := r.Group("/me")
	{
		m := MeController{}
		me.GET("", m.Index)
		me.POST("", m.UpdateMe)
		me.DELETE("", m.DeleteMe)
	}
}

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
	Setup(r)
	return router
}
