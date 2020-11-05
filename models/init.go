package models

import (
	"fmt"
	"time"

	"ffly-plus/internal/config"

	"github.com/colinrs/pkgx/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// DB ...
var DB *gorm.DB

// Database ...
func Database(mysqlConfig config.MySQLConfig) error {
	logger.Info("mysql {%#v}", mysqlConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Addr,
		mysqlConfig.DB)
	logger.Info("connect to mysql %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		logger.Error("conect db err:%v", err)
		return err
	}
	//设置连接池
	//空闲
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("get sqlDB err:%v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	//打开
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConfig.ConnMaxLifeTime))

	DB = db
	migration()
	return nil
}
