package main

import (
	_ "github.com/AlgernonGuo/tiktok-micro/internal/pkg/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	initRegister(h)
	h.Spin()
}
