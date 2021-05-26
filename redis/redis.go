package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type redisType struct {
	Pool *redis.Pool
}

var (
	redisClient *redisType
	once        sync.Once
)

const (
	RedisIP = "0.0.0.0"
)

func GetRedisConn() redis.Conn {
	once.Do(func() {
		redisPool := &redis.Pool{
			MaxActive: 100,
			Dial: func() (redis.Conn, error) {
				rc, err := redis.Dial("tcp", RedisIP+":6379")
				if err != nil {
					fmt.Println("Error connecting to redis:", err.Error())
					return nil, err
				}
				return rc, nil
			},
		}
		redisClient = &redisType{
			Pool: redisPool,
		}
	})
	return redisClient.Pool.Get()
}









