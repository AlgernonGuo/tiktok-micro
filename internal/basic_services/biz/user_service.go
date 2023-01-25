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

func (*UserService) GetUserById(targetUserId, myId int64) (user *data.User, err error) {
	user, err = data.NewUserDao().GetUserById(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetUserByName failed, err: %v", err)
		return nil, err
	}
	// hide password to make sure security
	user.Password = ""
	user.FollowCount, err = data.NewFollowDao().GetFollowCountByUserId(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetFollowCountByUserId failed, err: %v", err)
		return nil, err
	}
	user.FollowerCount, err = data.NewFollowDao().GetFollowerCountByFollowId(targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).Errorf("GetFollowerCountByFollowId failed, err: %v", err)
		return nil, err
	}
	if myId == targetUserId {
		user.IsFollow = true
		return user, nil
	}
	followInfo, err := data.NewFollowDao().GetFollowByUserIdAndFollowId(myId, targetUserId)
	if err != nil {
		logrus.WithField("targetUserId", targetUserId).WithField("userId", myId).Errorf("GetFollowByUserIdAndFollowId failed, err: %v", err)
		return nil, err
	}
	if followInfo.Id != 0 {
		user.IsFollow = true
	}
	return user, nil
}
