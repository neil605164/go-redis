package main

import (
	"go-redis/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.RouteProvider(r)

	r.Run()

}
