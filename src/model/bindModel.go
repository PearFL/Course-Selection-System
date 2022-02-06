package model

type Bind struct {
	TeacherID int `json:"teacher_id" form:"teacher_id" gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CourseID  int `json:"course_id" form:"course_id" gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
}

func (Bind) TableName() string {
	return "bind"
}
