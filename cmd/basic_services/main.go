package main

import (
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz/mw"
	_ "github.com/AlgernonGuo/tiktok-micro/internal/pkg/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

func main() {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
	)
	mw.InitJwt()
	InitRegister(h)
	h.Spin()
}
