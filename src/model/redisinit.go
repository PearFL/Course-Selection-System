package model

import (
	"course_select/src/database"
)

func init() {
	// 清空redis并将mysql中的表注入redis
	rc := database.RedisClient.Get()
	defer rc.Close()

	rc.Flush()

	var courses []Course
	db.Model(&Course{}).Find(&courses)
	for _, v := range courses {
		rc.Do("HSET", "CourseToCount", v.CourseID, v.Capacity-v.CapSelected)
		rc.Do("HSET", "CourseToName", v.CourseID, v.Name)
	}

	var binds []Bind
	db.Model(&Bind{}).Find(&binds)
	for _, v := range binds {
		rc.Do("HSET", "CourseToTeacher", v.CourseID, v.TeacherID)
	}
}
