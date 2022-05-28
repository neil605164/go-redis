package router

import (
	"go-redis/app/handler"

	"github.com/gin-gonic/gin"
)

func RouteProvider(r *gin.Engine) {
	r.GET("/", handler.Ping)
	r.GET("/set", handler.SetValue)
	r.GET("/get", handler.GetValue)
	r.GET("/stat", handler.PrintRedisPoolInfo)
}
