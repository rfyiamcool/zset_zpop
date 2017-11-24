package zpop

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	ErrRedisNotOpen = errors.New("ErrRedisNotOpen")

	ZpopScript = redis.NewScript(1, `
    local r = redis.call('ZRANGE', KEYS[1], 0, 0)
    if r ~= nil then
        r = r[1]
        redis.call('ZREM', KEYS[1], r)
    end
    return r
	`)
)

type RedisConfType struct {
	RedisPw          string
	RedisHost        string
	RedisDb          int
	RedisMaxActive   int
	RedisMaxIdle     int
	RedisIdleTimeOut int
}

func NewRedisPool(redis_conf RedisConfType) *redis.Pool {
	redis_client_pool := &redis.Pool{
		MaxIdle:     redis_conf.RedisMaxIdle,
		MaxActive:   redis_conf.RedisMaxActive,
		IdleTimeout: time.Duration(redis_conf.RedisIdleTimeOut) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_conf.RedisHost)
			if err != nil {
				return nil, err
			}

			// 选择db
			c.Do("SELECT", redis_conf.RedisDb)

			if redis_conf.RedisPw == "" {
				return c, nil
			}

			_, err = c.Do("AUTH", redis_conf.RedisPw)
			if err != nil {
				panic("redis password error")
			}

			return c, nil
		},
	}
	return redis_client_pool
}

func Zpop(rc redis.Conn, key string) (result string, err error) {
	result, err = redis.String(ZpopScript.Do(rc, "zz"))
	if err == redis.ErrNil {
		return result, nil
	}
	return result, nil
}

