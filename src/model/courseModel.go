package model

type Course struct {
	CourseID  int    `json:"course_id" form:"course_id" gorm:"primary_key"`
	TeacherID int    `json:"teacher_id" form:"teacher_id"`
	Name      string `json:"name" form:"name"`
	Capacity  int    `json:"capacity" form:"capacity"`
}

func (Course) TableName() string {
	return "course"
}
