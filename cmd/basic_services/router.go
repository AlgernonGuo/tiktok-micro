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

	// the service that does not need authentication
	noAuth := h.Group("/douyin")
	{
		login := noAuth.Group("/user")
		{
			login.POST("/register", func(c context.Context, ctx *app.RequestContext) {
				// if register success then auto login
				if err := service.Register(c, ctx); err != nil {
					ctx.JSON(http.StatusOK, mw.UserLoginResponse{
						Response: utils.Response{StatusCode: 400, StatusMsg: err.Error()},
					})
					return
				}
				mw.JwtMiddleware.LoginHandler(c, ctx)
			})
			login.POST("/login", mw.JwtMiddleware.LoginHandler)
		}
	}

	// the service that need authentication
	withAuth := h.Group("/douyin", mw.JwtMiddleware.MiddlewareFunc())
	{
		user := withAuth.Group("/user")
		{
			user.GET("/", service.GetUserInfo)
		}
	}
}
