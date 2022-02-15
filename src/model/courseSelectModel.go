package model

import "github.com/gomodule/redigo/redis"

func IncrAndGet(courseId string, rdb redis.Conn) {
	rdb.Do("HINCRBY", "CourseToCount", courseId, 1)
}

func DecrAndGet(courseId string, rdb redis.Conn) int {
	count, _ := redis.Int(rdb.Do("HINCRBY", "CourseToCount", courseId, -1))
	return count
}

func UpdateStudentCourse(studentId string, courseId string, rdb redis.Conn) {
	rdb.Do("SADD", studentId, courseId)
}

func GetStudentCourses(studentId string, rdb redis.Conn) []string {
	result, _ := redis.Strings(rdb.Do("SGET", studentId))
	return result
}

func GetCourseNameById(courseId string, rdb redis.Conn) string {
	result, _ := redis.String(rdb.Do("HGET", "CourseToName", courseId))
	return result
}

func GetTeacherByCourseId(courseId string, rdb redis.Conn) string {
	result, _ := redis.String(rdb.Do("HGET", "CourseToTeacher", courseId))
	return result
}

func IsBooked(studentId string, courseId string, rdb redis.Conn) bool {
	result, _ := redis.Bool(rdb.Do("SISMEMBER", studentId, courseId))
	return result
}
