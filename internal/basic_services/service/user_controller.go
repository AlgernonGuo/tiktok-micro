package service

import (
	"context"
	"net/http"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz"
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type UserResponse struct {
	utils.Response
	User data.User `json:"user"`
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	// get user id from context
	userId := c.GetInt64("identity")
	// get user info from db (already hide password)
	user, err := biz.NewUserService().GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "failed"},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: utils.Response{StatusCode: 200, StatusMsg: "success"},
		User:     *user,
	})
	return
}
