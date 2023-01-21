package data

import "time"

type Video struct {
	Id            int64     `json:"id,omitempty"`
	Author        User      `json:"author" ;gorm:"column:user_id"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty" ;gorm:"default:0"`
	CommentCount  int64     `json:"comment_count,omitempty" ;gorm:"default:0"`
	IsFavorite    bool      `json:"is_favorite,omitempty" ;gorm:"-"`
	CreatedAt     time.Time `json:"create_at,omitempty"`
}

func (Video) TableName() string {
	return "video"
}
