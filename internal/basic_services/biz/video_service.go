package biz

import (
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
)

// VideoService is the interface of video service.
type VideoService interface {
	// GetFeed returns the feed of the user.
	GetFeed(latestTime time.Time) ([]*data.Video, time.Time, error)
	// GetVideoListByUserId returns the video list of the user.
	GetVideoListByUserId(userId int64) ([]*data.Video, error)
}

// NewVideoService returns a new VideoService.
func NewVideoService() VideoService {
	return &videoService{data.NewVideoDao()}
}

type videoService struct {
	videoDao *data.VideoDao
}

func (s *videoService) GetFeed(latestTime time.Time) ([]*data.Video, time.Time, error) {
	feed, err := s.videoDao.GetFeed(latestTime)
	if err != nil {
		return nil, time.Time{}, err
	}
	// get now sec
	nextTime := time.Now()
	// get the latest time of the feed
	// the last one is the latest
	if len(feed) > 0 {
		nextTime = feed[len(feed)-1].CreatedAt
	}
	return feed, nextTime, nil
}

// GetVideoListByUserId Get video publish list by user id
func (s *videoService) GetVideoListByUserId(userId int64) ([]*data.Video, error) {
	return s.videoDao.GetVideoListByUserId(userId)
}
