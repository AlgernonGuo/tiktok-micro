package biz

import (
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	"github.com/sirupsen/logrus"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) GetUserById(userId int64) (user *data.User, err error) {
	user, err = data.NewUserDao().GetUserById(userId)
	if err != nil {
		logrus.WithField("username", userId).Errorf("GetUserByName failed, err: %v", err)
		return nil, err
	}
	// hide password to make sure security
	user.Password = ""
	return user, nil
}
