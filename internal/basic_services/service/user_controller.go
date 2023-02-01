package service

import (
	"context"
	"net/http"
	"strconv"

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
	// get targetUserId from query
	targetUserIdStr := c.Query("user_id")
	// string to int64
	targetUserId, err := strconv.ParseInt(targetUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "failed"},
		})
		return
	}
	// get user id from context
	userId := c.GetInt64("identity")
	// get user info from db (already hide password)
	user, err := biz.NewUserService().GetUserById(targetUserId, userId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "failed"},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: utils.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *user,
	})
}
