package main

import (
	"bubble/dao"
	"bubble/log"
	"bubble/models"
	"bubble/routers"
)


func main() {
	// 创建数据库
	// sql: creat database bubble;

	// 连接数据库  
	err := dao.InitMySQL()
	if err != nil {
		panic("连接数据库出错"+err.Error())
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	log.InitFile("log/logbydate", "bubble")
	router := routers.SetupRouter()
	router.Run(":8080")
}