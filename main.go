package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kylesliu/gin_exporter"
)

func main() {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	gin.IsDebugging()

	app.GET("demo1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "demo1",
		})
	})

	app.GET("demo2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "demo1",
		})
	})

	app.GET("demo3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "demo1",
		})
	})

	app = gin_exporter.Default(app)

	if err := app.Run("127.0.0.1:9000"); err != nil {
		panic(err.Error())
	}
}

func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		c.Next()
	}
}
