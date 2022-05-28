package router

import (
	"go-redis/app/handler"

	"github.com/gin-gonic/gin"
)

func RouteProvider(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	r.GET("/redis_conn", handler.RedisConnTest)
}
