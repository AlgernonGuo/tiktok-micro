package service

import (
	"context"
	"net/http"
	"strconv"
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
	latestTimeUnixStr := c.Query("latest_time")
	// if the latest time is "0", set it to the current time
	var latestTime time.Time
	if latestTimeUnixStr == "0" {
		latestTime = time.Now()
	} else {
		// convert the latest time to time.Time
		latestTimeUnix, err := strconv.ParseInt(latestTimeUnixStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: utils.Response{StatusCode: 400, StatusMsg: "the latest_time is invalid"},
			})
		}
		latestTime = time.Unix(latestTimeUnix, 0)
		if latestTime.Year() > 3000 || latestTime.Year() < 1970 {
			latestTime = time.Now()
		}
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
		Response: utils.Response{StatusCode: 0},
		Feed:     feed,
		NextTime: t.Unix(),
	})
}
