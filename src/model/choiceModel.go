package model

type Choice struct {
	StudentID int `json:"student_id" form:"student_id" gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CourseID  int `json:"course_id" form:"course_id" gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
}

func (Choice) TableName() string {
	return "choice"
}
