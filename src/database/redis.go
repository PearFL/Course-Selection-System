package database

import (
	"course_select/src/config"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisClient *redis.Pool

func init() {
	redisConf := config.GetRedisConfig()
	RedisClient = &redis.Pool{
		MaxIdle:     redisConf.MaxIdle,
		MaxActive:   redisConf.MaxActive,
		IdleTimeout: time.Duration(redisConf.TimeOut) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(redisConf.Type, redisConf.Redis_Host)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			/*if _, err := c.Do("AUTH", redisConf.AUTH); err != nil {
				_ = c.Close()
				fmt.Println(err.Error())
				return nil, err
			}*/
			return c, nil
		},
	}
}
