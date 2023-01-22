package biz

import (
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	log "github.com/sirupsen/logrus"
)

type LoginService struct {
	userDao *data.UserDao
}

func NewLoginService() *LoginService {
	return &LoginService{userDao: data.NewUserDao()}
}

func (*LoginService) Login(username, password string) (user *data.User, err error) {
	user, err = data.NewUserDao().GetUserByNameAndPassword(username, password)
	if err != nil {
		log.WithField("username", username).Errorf("GetUserByName failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (*LoginService) Register(user *data.User) error {
	err := data.NewUserDao().CreateUser(user)
	if err != nil {
		log.WithField("username", user.Name).Warnf("CreateUser failed, err: %v", err)
		return err
	}
	return nil
}
