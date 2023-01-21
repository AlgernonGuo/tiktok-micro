package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
)

var (
	logPath = "./"
	logFile = time.Now().Format("2006-01-02") + ".log"
)

func initHlog() error {
	// TODO read config file to set log default output

	// set log output to file and console
	file, err := os.OpenFile(logPath+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logrus.Warning("Failed to log to file, using default stderr")
	} else {
		logrus.SetOutput(io.MultiWriter(file, os.Stdout))
		hlog.SetOutput(io.MultiWriter(file, os.Stdout))
	}

	logger := hertzlogrus.NewLogger()
	logger.Logger().SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
		ForceColors:     true,
		DisableColors:   false,
	})
	hlog.SetLogger(logger)
	return nil
}

func initLogrus() error {

	// TODO read config file to set log default output

	// set log output to file and console
	file, err := os.OpenFile(logPath+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logrus.Warning("Failed to log to file, using default stderr")
	} else {
		logrus.SetOutput(io.MultiWriter(file, os.Stdout))
		hlog.SetOutput(io.MultiWriter(file, os.Stdout))
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
		ForceColors:     true,
		DisableColors:   false,
	})
	return nil
}

func init() {
	err := initHlog()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = initLogrus()
	if err != nil {
		fmt.Println(err)
		return
	}
	logrus.Infof("logrus init success")
	hlog.Infof("hlog init success")
}
