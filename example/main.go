package main

import (
	"fmt"

	zpop "github.com/rfyiamcool/zset_zpop"
)

func main() {

	key := "zz"

	redisConfig := zpop.RedisConfType{
		RedisPw:          "",
		RedisHost:        "127.0.0.1:6379",
		RedisDb:          0,
		RedisMaxActive:   100,
		RedisMaxIdle:     100,
		RedisIdleTimeOut: 1000,
	}

	redisPool := zpop.NewRedisPool(redisConfig)

	rc := redisPool.Get()
	defer rc.Close()

	v, err := zpop.Zpop(rc, key)
	fmt.Printf("value: %s, error: %v", v, err)
}
