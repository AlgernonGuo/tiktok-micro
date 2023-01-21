package mysql

import (
	"fmt"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	username := "root"
	password := "8520"
	host := "127.0.0.1"
	port := 3306
	Dbname := "tiktok_micro"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,                                // able to use Prepare Statement to improve performance
		SkipDefaultTransaction: true,                                // disable default transaction
		Logger:                 logger.Default.LogMode(logger.Info), // set log level
	})
	if err != nil {
		log.WithError(err).Panic("failed to connect database")
	}

	sqlDB, err := _db.DB()
	if err != nil {
		log.WithError(err).Panic("failed to connect database")
	}
	// Set max idle connections 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// Set max open connections 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	logrus.Infof("database init success")
}

func GetDB() *gorm.DB {
	return _db
}
