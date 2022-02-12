package model

import (
	"course_select/src/database"
	"errors"
)

type Bind struct {
	TeacherID int `json:"teacher_id" form:"teacher_id" gorm:"primary_key"`
	CourseID  int `json:"course_id" form:"course_id" gorm:"primary_key"`
}

func (Bind) TableName() string {
	return "bind"
}

func BindCourse(bind Bind) error {
	//go get -u gorm.io/gorm // get gorm v2
	//go get -u gorm.io/driver/mysql // get dialector of mysql from gorm
	//db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&user)
	//var db := database.MySqlDb.DB()
	result := database.MySqlDb.Exec("INSERT IGNORE INTO bind(teacher_id,course_id) VALUES (?,?)", bind.TeacherID, bind.CourseID)
	if result.RowsAffected == 0 {
		return errors.New("CourseHasBound")
	}
	return nil
}

func UnBindCourse(unbind Bind) error {
	result := database.MySqlDb.Delete(&unbind)
	if result.RowsAffected == 0 {
		return errors.New("CourseNotBind")
	}
	return nil
}
