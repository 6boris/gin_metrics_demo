package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gin_metrics "github.com/kylesliu/gin_exporter"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	app := gin.New()
	gin.SetMode(gin.DebugMode)
	gin_metrics.Default(app)

	app.GET("rand_str", func(c *gin.Context) {
		c.JSON(200, getRandomString(32))
	})
	app.GET("rand_int", func(c *gin.Context) {
		c.JSON(200, rand.Intn(100000000000))
	})
	app.GET("time_now", func(c *gin.Context) {
		c.JSON(200, time.Now())
	})
	app.GET("rand_sleep", func(c *gin.Context) {
		time.Sleep(time.Duration(time.Duration(rand.Intn(1000)) * time.Millisecond))
		c.JSON(200, time.Now())
	})
	for i := 0; i < 100; i++ {
		app.GET(fmt.Sprint("rand_route", strconv.Itoa(i)), func(c *gin.Context) {
			c.JSON(200, time.Now())
		})
	}

	if err := app.Run("127.0.0.1:9000"); err != nil {
		panic(err.Error())
	}
}

func getRandomString(str_len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < str_len; i++ {
		result = append(result, bytes[r.Intn(str_len)])
	}
	return string(result)
}
