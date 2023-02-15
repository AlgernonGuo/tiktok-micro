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
	if latestTimeUnixStr == "0" || latestTimeUnixStr == "" {
		latestTime = time.Now()
	} else {
		// convert the latest time to time.Time
		latestTimeUnix, err := strconv.ParseInt(latestTimeUnixStr, 10, 64)
		if err != nil {
			logrus.WithField("latest_time", latestTimeUnixStr).WithField("latest_time_unix", latestTimeUnix).Infof("GetFeed failed, err: %v", err)
			c.JSON(http.StatusOK, FeedResponse{
				Response: utils.Response{StatusCode: 400, StatusMsg: "the latest_time is invalid"},
			})
			return
		}
		latestTime = time.Unix(latestTimeUnix, 0)
		if latestTime.Year() > 3000 || latestTime.Year() < 1970 {
			latestTime = time.Now()
		}
	}
	// get feed from db
	feed, t, err := biz.NewVideoService().GetFeed(latestTime)
	if err != nil {
		logrus.WithField("latest_time", latestTime).Warnf("GetFeed failed, err: %v", err)
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

// GetPublishVideoList Get video list of a user
func GetPublishVideoList(ctx context.Context, c *app.RequestContext) {
	// get user id from context
	userIdStr := c.Query("user_id")
	// string to int64
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logrus.WithField("user_id", userIdStr).Warnf("GetPublishVideoList failed, err: %v", err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "user_id is invalid"},
		})
		return
	}
	// get video list from db
	videoList, err := biz.NewVideoService().GetVideoListByUserId(userId)
	if err != nil {
		logrus.WithField("user_id", userId).Errorf("GetVideoListByUserId failed, err: %v", err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: utils.Response{StatusCode: 400, StatusMsg: "server error"},
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: utils.Response{StatusCode: 0},
		Feed:     videoList,
	})
}

// upload video
func UploadVideo(ctx context.Context, c *app.RequestContext) {
	userId := c.GetInt64("identity")
	// get video title from body form
	title := c.FormValue("title")
	file, err := c.FormFile("data")
	if err != nil {
		logrus.WithField("user_id", userId).Errorf("UploadVideo failed, err: %v", err)
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 400, StatusMsg: "failed",
		})
		return
	}

	// file name is user id + random string
	// user id 前5位 + 随机字符串
	fileName := strconv.FormatInt(userId%1000, 10) + utils.GetRandomString(5)
	err = biz.NewVideoService().UploadVideo(ctx, userId, string(title), fileName, file)
	if err != nil {
		logrus.WithField("user_id", userId).Errorf("UploadVideo failed, err: %v", err)
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 400, StatusMsg: "failed",
		})
		return
	}
	videoList, err := biz.NewVideoService().GetVideoListByUserId(userId)
	if err != nil {
		logrus.WithField("user_id", userId).Errorf("GetVideoListByUserId failed, err: %v", err)
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 400, StatusMsg: "server error",
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: utils.Response{StatusCode: 0},
		Feed:     videoList,
	})
}
