package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func initLogrus() error {

	// TODO 读取配置文件设置log默认文件输出
	// log.Formatter = &logrus.JSONFormatter{}
	//file, err := os.OpenFile("./gin_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Error("Failed to log to file, using default stderr")
	//	return err
	//}
	//log.Out = file
	log.Out = os.Stdout

	// set gin default log to file
	// gin.DefaultWriter = file
	// gin.SetMode(gin.ReleaseMode)

	log.Level = logrus.InfoLevel
	return nil
}

func init() {
	err := initLogrus()
	if err != nil {
		fmt.Println(err)
		return
	}
	logrus.Infof("logrus init success")
}
