package models

import (
	"fmt"
	"time"

	"ffly-plus/internal/config"

	"github.com/colinrs/pkgx/logger"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database ...
func Database(mysql config.MySQLConfig) error {
	logger.Info("mysql {%#v}", mysql)
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.UserName,
		mysql.Password,
		mysql.Addr,
		mysql.DB)
	logger.Info("connect to mysql %s", connString)
	db, err := gorm.Open("mysql", connString)

	db.LogMode(true)
	// Error
	if err != nil {
		logger.Error("conect db err:%v", err)
		return err
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(mysql.MaxIdleConn)
	//打开
	db.DB().SetMaxOpenConns(mysql.MaxOpenConn)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(mysql.ConnMaxLifeTime))

	DB = db
	return nil
}
