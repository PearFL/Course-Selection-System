package model

import "github.com/gomodule/redigo/redis"

const (
	StudentPrefix string = "sid:"
	TeacherPrefix string = "tid:"
	CoursePrefix  string = "cid:"
)

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

func TeacherBindCourse(teacherId string, courseId string, rdb redis.Conn) {
	rdb.Do("SADD", TeacherPrefix+teacherId, CoursePrefix+courseId)
}

func TeacherUnbindCourse(teacherId string, rdb redis.Conn) {
	rdb.Do("SDEL", TeacherPrefix+teacherId)
}
