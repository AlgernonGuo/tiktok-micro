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

func (*LoginService) Login(username, password string) (token string, user *data.User, err error) {
	user, err = data.NewUserDao().GetUserByName(username)
	if err != nil {
		log.WithField("username", username).Errorf("GetUserByName failed, err: %v", err)
		return "", nil, err
	}
	token = user.Name + user.Password // TODO: use JWT
	return token, user, nil
}

func (*LoginService) Register(user *data.User) (token string, err error) {
	err = data.NewUserDao().CreateUser(user)
	if err != nil {
		log.WithField("username", user.Name).Errorf("CreateUser failed, err: %v", err)
		return "", err
	}
	token = user.Name + user.Password // TODO: use JWT
	return token, nil
}
