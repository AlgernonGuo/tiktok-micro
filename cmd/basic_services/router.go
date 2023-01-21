package main

import (
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func initRegister(h *server.Hertz) {
	//h.StaticFS("/static", &app.FS{Root: "./", GenerateIndexPages: true})
	h.Static("/static", "./web/static")
	RegisterGroupRoute(h)
}

// RegisterGroupRoute group route
func RegisterGroupRoute(h *server.Hertz) {
	// User group:
	user := h.Group("/douyin/user")
	{
		// loginEndpoint is a handler func
		user.POST("/register", service.Register)
		user.POST("/login", service.Login)
	}
}
