package repository

import (
	"github.com/AlgernonGuo/tiktok-micro/utils"
	"gorm.io/gorm"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	CreateTime    string `json:"create_time,omitempty"`
}

func (User) TableName() string {
	return "user"
}

// UserDao
type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// GetUserById
func (u *UserDao) GetUserById(id int64) (*User, error) {
	// TODO
	return nil, nil
}

// GetUserByName
func (u *UserDao) GetUserByName(username string) (user *User, err error) {
	db := utils.GetDB()
	if db == nil {
		return nil, nil
	}
	err = db.Table("user").Where("name = ?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return user, err
}
