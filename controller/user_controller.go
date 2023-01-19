package controller

import (
	"github.com/AlgernonGuo/tiktok-micro/repository"
	"github.com/AlgernonGuo/tiktok-micro/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserLoginResponse struct {
	utils.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	utils.Response
	User repository.User `json:"user"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, err := repository.NewUserDao().GetUserByName(username)
	if err != nil {
		logrus.Infof("Get user by name error: %v", err)
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User not exist"},
		})
		return
	}

	token := username + password // TODO: use JWT
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})

}
