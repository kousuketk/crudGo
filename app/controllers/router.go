package controllers

import (
	"github.com/gin-gonic/gin"
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
