package main

import (
	"github.com/AlgernonGuo/tiktok-micro/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./resources/static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/login/", controller.Login)
}
