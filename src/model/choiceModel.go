package model

type Choice struct {
	StudentID string `json:"student_id" form:"student_id" gorm:"primary_key"`
	CourseID  string `json:"course_id" form:"course_id" gorm:"primary_key"`
}

func (Choice) TableName() string {
	return "choice"
}
