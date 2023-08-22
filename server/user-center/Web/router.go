package main

import (
	"github.com/gin-gonic/gin"

	"web/api"
	"web/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	g := r.Group("/douyin")
	{
		g.POST("user/register/", api.UserRegister)
		g.POST("user/login/", api.UserLogin)
		authed := g.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			authed.GET("user/", api.GetUserById)
			authed.POST("/publish/action/")
			authed.GET("/publish/list/")

			// extra apis - I
			authed.POST("/favorite/action/")
			authed.GET("/favorite/list/")
			authed.POST("/comment/action/")
			authed.GET("/comment/list/")

			// extra apis - II
			authed.POST("/relation/action/")
			authed.GET("/relation/follow/list/")
			authed.GET("/relation/follower/list/")
			authed.GET("/relation/friend/list/")
			authed.GET("/message/chat/")
			authed.POST("/message/action/")
		}

	}
	return r
}
