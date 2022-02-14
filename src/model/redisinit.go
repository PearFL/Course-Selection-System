package model

import (
	"course_select/src/database"
)

func init() {
	// 清空redis并将mysql中的表注入redis
	get := database.RedisClient.Get()
	get.Flush()

	var courses []Course
	db.Model(&Course{}).Find(&courses)
	for _, v := range courses {
		get.Do("HSET", "CourseToCount", v.CourseID, v.Capacity-v.CapSelected)
		get.Do("HSET", "CourseToName", v.CourseID, v.Name)
	}

	var binds []Bind
	db.Model(&Bind{}).Find(&binds)
	for _, v := range binds {
		get.Do("HSET", "CourseToTeacher", v.CourseID, v.TeacherID)
	}
}
