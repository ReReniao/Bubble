package routers

import (
	"bubble/controller"
	"bubble/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter()  *gin.Engine {
	// router := gin.Default()

	// 使用自己的日志
	// 使用New创建路由
	router := gin.New()
	// 创建日志中间件
	router.Use(middleware.LogMiddleware())

	// 告诉gin框架模版文件引用的静态文件去哪里找
	router.Static("/static", "static")

	// 告诉gin框架区哪里找模版文件
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)
	// v1
	v1Group := router.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看一个待办事项
		v1Group.GET("todo/:id", func(c *gin.Context) {

		})

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}
	return router
}