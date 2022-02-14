package database

import (
	"course_select/src/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var RedisClient *redis.Pool

const (
	StudentPrefix string = "sid:"
	TeacherPrefix string = "tid:"
	CoursePrefix  string = "cid:"
)

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
}

func IncrAndGet(courseId string, rdb redis.Conn) {
	rdb.Do("HINCRBY CourseToCount " + courseId + " 1")
}

func DecrAndGet(courseId string, rdb redis.Conn) int {
	count, _ := redis.Int(rdb.Do("HINCRBY", "CourseToCount", CoursePrefix+courseId, -1))
	return count
}

func UpdateStudentCourse(studentId string, courseId string, rdb redis.Conn) {
	rdb.Do("SADD", StudentPrefix+studentId, CoursePrefix+courseId)
}

func GetStudentCourses(studentId string, rdb redis.Conn) []string {
	result, _ := redis.Strings(rdb.Do("SGET", StudentPrefix+studentId))
	return result
}

func UpdateTeacherCourse(teacherId string, courseId string, rdb redis.Conn) {
	rdb.Do("SADD", TeacherPrefix+teacherId, CoursePrefix+courseId)
}
