package model

type Bind struct {
	TeacherID string `json:"teacher_id" form:"teacher_id" gorm:"primary_key"`
	CourseID  string `json:"course_id" form:"course_id" gorm:"primary_key"`
}

func (Bind) TableName() string {
	return "bind"
}
