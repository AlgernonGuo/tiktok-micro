package repository

import "time"

type Favorite struct {
	Id        int64     `json:"id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	VideoId   int64     `json:"video_id,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	Status    int8      `json:"status,omitempty" ;gorm:"default:1"`
}

func (Favorite) TableName() string {
	return "favorite"
}
