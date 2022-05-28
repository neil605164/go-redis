package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var gClient *redis.Client

func RedisConn() *redis.Client {
	if gClient != nil {
		fmt.Println("use old====>", gClient)
		return gClient
	}

	fmt.Println("connect new====>", gClient)

	gClient = redis.NewClient(&redis.Options{
		//连接信息
		Addr:     "redis_go-redis:6379", //主机名+冒号+端口，默认localhost:6379
		Password: "rd3Redis",            //密码

		//连接池容量及闲置连接数量
		PoolSize:    2,                // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		IdleTimeout: 20 * time.Second, //闲置超时，默认5分钟，-1表示取消闲置超时检查

	})

	return gClient
}

func PrintRedisPool(stats *redis.PoolStats) {
	fmt.Printf("Hits=%d Misses=%d Timeouts=%d TotalConns=%d IdleConns=%d StaleConns=%d\n",
		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns)
}

func PrintRedisOption(opt *redis.Options) {
	fmt.Printf("Network=%v\n", opt.Network)
	fmt.Printf("Addr=%v\n", opt.Addr)
	fmt.Printf("Password=%v\n", opt.Password)
	fmt.Printf("DB=%v\n", opt.DB)
	fmt.Printf("MaxRetries=%v\n", opt.MaxRetries)
	fmt.Printf("MinRetryBackoff=%v\n", opt.MinRetryBackoff)
	fmt.Printf("MaxRetryBackoff=%v\n", opt.MaxRetryBackoff)
	fmt.Printf("DialTimeout=%v\n", opt.DialTimeout)
	fmt.Printf("ReadTimeout=%v\n", opt.ReadTimeout)
	fmt.Printf("WriteTimeout=%v\n", opt.WriteTimeout)
	fmt.Printf("PoolSize=%v\n", opt.PoolSize)
	fmt.Printf("MinIdleConns=%v\n", opt.MinIdleConns)
	fmt.Printf("MaxConnAge=%v\n", opt.MaxConnAge)
	fmt.Printf("PoolTimeout=%v\n", opt.PoolTimeout)
	fmt.Printf("IdleTimeout=%v\n", opt.IdleTimeout)
	fmt.Printf("IdleCheckFrequency=%v\n", opt.IdleCheckFrequency)
	fmt.Printf("TLSConfig=%v\n", opt.TLSConfig)

}
