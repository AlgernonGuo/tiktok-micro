package service

import (
	"context"
	"net/http"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz"
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	log "github.com/sirupsen/logrus"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserLoginResponse struct {
	utils.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	utils.Response
	User data.User `json:"user"`
}

func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	token, user, err := biz.NewLoginService().Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User not exist"},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})
}

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	user := &data.User{
		Id:       utils.GenID(),
		Name:     username,
		Password: password,
	}
	token, err := biz.NewLoginService().Register(user)
	if err != nil {
		log.WithField("username", username).Errorf("GetUserByName failed, err: %v", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})
}
