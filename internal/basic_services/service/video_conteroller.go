package service

import (
	"context"
	"net/http"
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz"
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
)

type FeedResponse struct {
	utils.Response
	Feed     []*data.Video `json:"video_list,omitempty"`
	NextTime int64         `json:"next_time,omitempty"`
}

func GetFeed(ctx context.Context, c *app.RequestContext) {
	// get the latest time from the request query
	latestTimeStr := c.Query("latest_time")
	// string to time
	latestTime, err := time.Parse("2006-01-02 15:04:05", latestTimeStr)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "the latest_time is invalid"},
		})
		return
	}
	// get feed from db
	feed, t, err := biz.NewVideoService().GetFeed(latestTime)
	if err != nil {
		logrus.WithField("latest_time", latestTime).Errorf("GetFeed failed, err: %v", err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "failed"},
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: utils.Response{StatusCode: 200, StatusMsg: "success"},
		Feed:     feed,
		NextTime: t.Unix(),
	})
}
