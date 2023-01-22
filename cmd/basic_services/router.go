package main

import (
	"context"
	"net/http"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz/mw"
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/service"
	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRegister(h *server.Hertz) {
	//h.StaticFS("/static", &app.FS{Root: "./", GenerateIndexPages: true})
	h.Static("/static", "./web/static")
	RegisterGroupRoute(h)
}

// RegisterGroupRoute group route
func RegisterGroupRoute(h *server.Hertz) {
	root := h.Group("/douyin")
	{
		// User group:
		user := root.Group("/user")
		{
			user.POST("/register", func(c context.Context, ctx *app.RequestContext) {
				// if register success then auto login
				if err := service.Register(c, ctx); err != nil {
					ctx.JSON(http.StatusOK, mw.UserLoginResponse{
						Response: utils.Response{StatusCode: 400, StatusMsg: err.Error()},
					})
					return
				}
				mw.JwtMiddleware.LoginHandler(c, ctx)
			})
			user.POST("/login", mw.JwtMiddleware.LoginHandler)
		}
	}
}
