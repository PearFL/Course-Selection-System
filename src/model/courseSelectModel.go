package model

import "github.com/gomodule/redigo/redis"

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
