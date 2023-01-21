package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

var node *snowflake.Node

func init() {
	var startTime = "2023-01-01"
	var machineID = int64(1)
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		logrus.WithField("startTime", startTime).Errorf("Parse time failed, err: %v", err)
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err == nil {
		// log start time and machine id
		logrus.WithField("startTime", startTime).WithField("machineID", machineID).Info("Snowflake init success")
		return
	}
	logrus.WithField("machineID", machineID).Errorf("NewNode failed, err: %v", err)
	logrus.Errorf("Try to use machineID 1")
	node, err = snowflake.NewNode(1)
	if err != nil {
		logrus.Panic("Use machineID 1 to newNode failed")
		return
	}
}

// GenID generate a unique snowflake ID
func GenID() int64 {
	return node.Generate().Int64()
}
