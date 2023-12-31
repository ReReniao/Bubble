package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	status200 = 42
	status404 = 43
	status500 = 41

	methodGET = 44
	methodPUT = 45
	methodPOST = 46
	methodDELETE = 47
)

// 返回一个中间件函数
// 每次方法路由都会运行
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped

		// Stop timer
		end := time.Now()
		timeSub := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		//bodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		var statusColor string 
		switch statusCode {
		case 200:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status200, statusCode)
		case 404:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status404, statusCode)
		}

		// 一种包含颜色编码的字符串
		var methodColor string

		switch method {
		case "GET":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodGET, method)
		case "POST":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodPOST, method)
		case "PUT":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodPUT, method)
		case "DELETE":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodDELETE, method)
		}

		// logrus设置了保存到文件的hook以及日志等级
		logrus.Infof("[GIN] %s | %s | %d | %s | %s | %s ",
			start.Format("2006-01-02 15-04-05"),
			statusColor,
			timeSub,
			clientIP,
			methodColor,
			path,
		)
	}
	
}