package data

import (
	"errors"
	"time"

	sd "gorm.io/plugin/soft_delete"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/mysql"
)

type User struct {
	Id            int64        `json:"id,omitempty"`
	Name          string       `json:"name,omitempty"`
	Password      string       `json:"password,omitempty"`
	FollowCount   int64        `json:"follow_count,omitempty" ;gorm:"default:0"`
	FollowerCount int64        `json:"follower_count,omitempty" ;gorm:"default:0"`
	CreatedAt     time.Time    `json:"create_at,omitempty"`
	isDel         sd.DeletedAt `gorm:"softDelete:flag"`
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
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	err = db.Where("name = ?", username).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *UserDao) CreateUser(user *User) error {
	db := mysql.GetDB()
	if db == nil {
		return nil
	}
	// check the username if exist
	var count int64
	err := db.Model(&User{}).Where("name = ?", user.Name).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exist")
	}
	return db.Create(user).Error
}
