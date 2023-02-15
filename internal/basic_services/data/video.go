package data

import (
	"errors"
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/mysql"
	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
)

type Video struct {
	Id            int64     `json:"id,omitempty" gorm:"primary_key"`
	UserId        int64     `gorm:"column:user_id" json:"-"`
	Author        User      `gorm:"foreignKey:Id;references:user_id" json:"author"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	Title         string    `json:"title,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty" gorm:"default:0"`
	CommentCount  int64     `json:"comment_count,omitempty" gorm:"default:0"`
	IsFavorite    bool      `json:"is_favorite,omitempty" gorm:"-"`
	IsDel         bool      `gorm:"softDelete:flag" json:"-"`
	CreatedAt     time.Time `json:"create_at,omitempty" gorm:"autoCreateTime"`
}

func (Video) TableName() string {
	return "video"
}

// VideoDao
type VideoDao struct {
}

func NewVideoDao() *VideoDao {
	return &VideoDao{}
}

// GetFeed
func (v *VideoDao) GetFeed(latestTime time.Time) ([]*Video, error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, errors.New("db is nil")
	}
	var videos []*Video
	// limit 30
	err := db.Model(&Video{}).Preload("Author").Where("created_at < ?", latestTime).Order("id desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// GetVideoByUserId
func (v *VideoDao) GetVideoListByUserId(userId int64) ([]*Video, error) {
	db := mysql.GetDB()
	if db == nil {
		return nil, errors.New("db is nil")
	}
	var videos []*Video
	err := db.Model(&Video{}).Preload("Author").Where("user_id = ?", userId).Order("id desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *VideoDao) SaveVideo(userId int64, title string, name string) {
	db := mysql.GetDB()
	if db == nil {
		return
	}
	video := Video{
		Id:       utils.GenID(),
		Title:    title,
		UserId:   userId,
		PlayUrl:  utils.TencentCosURL + "/video/" + name + ".mp4",
		CoverUrl: utils.TencentCosURL + "/cover/" + name + ".jpg",
	}
	db.Model(&video).Create(&video)
}
