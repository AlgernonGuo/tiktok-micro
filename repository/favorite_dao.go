package repository

type Favorite struct {
	Id         int64  `json:"id,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	VideoId    int64  `json:"video_id,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Status     int8   `json:"status,omitempty"`
}

func (Favorite) TableName() string {
	return "favorite"
}
