package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Bubblelog logger.Interface

func InitMySQL() (err error) {
	username := "root"
	password := "xmhcp1472580369"
	host := "127.0.0.1"
	port := 3306
	DBname := "bubble"
	timeout := "10s"

	Bubblelog =  logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)

	// 连接数据库实例
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: Bubblelog,
	})

	if err != nil {
		return err
	}

	return nil

}
