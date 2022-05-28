package main

import (
	"go-redis/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", handler.Ping)
	r.GET("/set", handler.SetValue)
	r.GET("/get", handler.GetValue)
	r.GET("/stat", handler.PrintRedisPoolInfo)

	r.Run()

}
