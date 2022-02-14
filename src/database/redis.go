package database

import (
	"course_select/src/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
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
				log.Println(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", redisConf.AUTH); err != nil {
				_ = c.Close()
				log.Println(err.Error())
				return nil, err
			}
			return c, nil
		},
	}
	// TODO
	// 清空redis并将mysql中的表注入redis
	get := RedisClient.Get()
	get.Flush()
	/*find := model.GetAllCourse()
	fmt.Println(find)*/
	// get.Do("SET", CourseId, CourseRemainCap)

}
