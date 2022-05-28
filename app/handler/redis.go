package handler

import (
	"fmt"
	"go-redis/app/cache"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	gClient := cache.RedisConn()
	for i := 0; i < 10; i++ {
		gClient.Ping().Result()
	}
	// printRedisPool(gClient.PoolStats())
	c.JSON(http.StatusOK, gin.H{
		"time": time.Now(),
	})
}

func SetValue(c *gin.Context) {
	gClient := cache.RedisConn()
	err := gClient.Set("key", "value", 0).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
	if err != nil {
		panic(err)
	}

}

func GetValue(c *gin.Context) {
	gClient := cache.RedisConn()
	var val string
	var err error
	for i := 0; i < 10; i++ {
		val, err = gClient.Get("key").Result() // => GET key
		if err != nil {
			panic(err)
		}
		fmt.Println("key %v", i, val)
	}

	c.JSON(http.StatusOK, gin.H{
		"key": val,
	})

}

func PrintRedisPoolInfo(c *gin.Context) {

	gClient := cache.RedisConn()
	stats := gClient.PoolStats()
	str := fmt.Sprintf("Hits=%d Misses=%d Timeouts=%d TotalConns=%d IdleConns=%d StaleConns=%d\n",
		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns)

	c.JSON(http.StatusOK, gin.H{
		"result": str,
	})

}
