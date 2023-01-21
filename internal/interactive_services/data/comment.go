package repository

import (
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
)

type Comment struct {
	Id        int64     `json:"id,omitempty"`
	User      data.User `json:"user"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
}

func (Comment) TableName() string {
	return "comment"
}
