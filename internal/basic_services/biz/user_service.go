package biz

import (
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userDao   *data.UserDao
	followDao *data.FollowDao
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserById(targetUserId, myId int64) (user *data.User, err error) {
	user, err = s.userDao.GetUserById(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetUserByName failed, err: %v", err)
		return nil, err
	}
	// hide password to make sure security
	user.Password = ""
	user.FollowCount, err = s.followDao.GetFollowCountByUserId(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetFollowCountByUserId failed, err: %v", err)
		return nil, err
	}
	user.FollowerCount, err = s.followDao.GetFollowerCountByFollowId(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetFollowerCountByFollowId failed, err: %v", err)
		return nil, err
	}
	if myId == targetUserId {
		user.IsFollow = true
		return user, nil
	}
	followInfo, err := s.followDao.GetFollowByUserIdAndFollowId(myId, targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).WithField("userId", myId).Errorf("GetFollowByUserIdAndFollowId failed, err: %v", err)
		return nil, err
	}
	if followInfo.Id != 0 {
		user.IsFollow = true
	}
	return user, nil
}
