package router

import (
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	// 创建Gin引擎
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	return r
}
