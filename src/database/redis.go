package database

import (
	"course_select/src/config"
	global "course_select/src/global"
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
}

func IncrAndGet(courseId string, rdb redis.Conn) {
	rdb.Do("HINCRBY CourseToCount " + courseId + " 1")
}

func DecrAndGet(courseId string, rdb redis.Conn) int {
	count, _ := redis.Int(rdb.Do("HINCRBY", "CourseToCount", global.CourseIdPrefix+courseId, -1))
	return count
}

func UpdateStudentCourse(studentId string, courseId string, rdb redis.Conn) {
	rdb.Do("SADD", global.StudentIdPrefix+studentId, global.CourseIdPrefix+courseId)
}

func GetStudentCourses(studentId string, rdb redis.Conn) []string {
	result, _ := redis.Strings(rdb.Do("SGET", global.StudentIdPrefix+studentId))
	return result
}
