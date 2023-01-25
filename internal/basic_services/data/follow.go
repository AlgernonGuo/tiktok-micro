package data

import (
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/mysql"
)

type Follow struct {
	Id        int64 `json:"id,omitempty"`
	UserId    int64 `json:"user_id,omitempty"`   // 关注者id
	FollowId  int64 `json:"follow_id,omitempty"` // 被关注者id
	CreatedAt int64 `json:"created_at,omitempty"`
	isDel     int64 `gorm:"softDelete:flag"`
}

func (Follow) TableName() string {
	return "follow"
}

// FollowDao
type FollowDao struct {
}

func NewFollowDao() *FollowDao {
	return &FollowDao{}
}

// GetFollowByUserIdAndFollowId
func (f *FollowDao) GetFollowByUserIdAndFollowId(userId, followId int64) (*Follow, error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	var follow Follow
	err := db.Where("user_id = ? and follow_id = ?", userId, followId).Limit(1).Find(&follow).Error
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

// GetFollowListByUserId
func (f *FollowDao) GetFollowListByUserId(userId int64) ([]*Follow, error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	var follow []*Follow
	err := db.Where("user_id = ?", userId).Find(&follow).Error
	if err != nil {
		return nil, err
	}
	return follow, nil
}

// GetFollowerListByFollowId
func (f *FollowDao) GetFollowerListByFollowId(followId int64) ([]*Follow, error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, nil
	}
	var follow []*Follow
	err := db.Where("follow_id = ?", followId).Find(&follow).Error
	if err != nil {
		return nil, err
	}
	return follow, nil
}

// AddFollow
func (f *FollowDao) AddFollow(userId, followId int64) error {
	db := mysql.GetDB()
	if db == nil {
		return nil
	}
	follow := &Follow{
		UserId:    userId,
		FollowId:  followId,
		CreatedAt: time.Now().Unix(),
	}
	err := db.Create(follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *FollowDao) DelFollow(userId, followId int64) error {
	db := mysql.GetDB()
	if db == nil {
		return nil
	}
	err := db.Where("user_id = ? and follow_id = ?", userId, followId).Delete(&Follow{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetFollowCountByUserId 获取关注数
func (f *FollowDao) GetFollowCountByUserId(userId int64) (int64, error) {
	db := mysql.GetDB()
	if db == nil {
		return 0, nil
	}
	var count int64
	err := db.Model(&Follow{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFollowerCountByFollowId 获取粉丝数
func (f *FollowDao) GetFollowerCountByFollowId(followId int64) (int64, error) {
	db := mysql.GetDB()
	if db == nil {
		return 0, nil
	}
	var count int64
	err := db.Model(&Follow{}).Where("follow_id = ?", followId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
