package data

import (
	"errors"
	"time"

	"gorm.io/gorm"

	sd "gorm.io/plugin/soft_delete"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/mysql"
)

type User struct {
	Id            int64        `json:"id,omitempty"`
	Name          string       `json:"name,omitempty"`
	Password      string       `json:"password,omitempty"`
	FollowCount   int64        `json:"follow_count" gorm:"default:0"`
	FollowerCount int64        `json:"follower_count" gorm:"default:0"`
	CreatedAt     time.Time    `json:"-"`
	IsFollow      bool         `json:"is_follow" gorm:"-"`
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
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	var user User
	err := db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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

func (u *UserDao) GetUserByNameAndPassword(username, password string) (user *User, err error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	err = db.Where("name = ? and password = ?", username, password).Take(&user).Error
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
	// user transaction
	if err := db.Transaction(func(tx *gorm.DB) error {
		// check the username if exist
		var count int64
		if err := tx.Model(&User{}).Where("name = ?", user.Name).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("username already exist")
		}
		err := db.Create(user).Error
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
