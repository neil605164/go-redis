package main

import (
	"fmt"
	"go-redis/router"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("%d CPU burnner running ...\n", runtime.NumCPU())
	r := gin.Default()

	router.RouteProvider(r)

	r.Run()

}
