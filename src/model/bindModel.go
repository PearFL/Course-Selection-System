package model

import (
	global "course_select/src/global"
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
	var member Member
	err := db.First(&Member{}, "user_id = ?", bind.TeacherID).Scan(&member).Error
	if err != nil || member.IsDeleted == true || member.UserType != global.Teacher {
		return errors.New("TeacherNotExisted")
	}

	var course Course
	err = db.Model(&Course{}).Where("course_id = ?", bind.CourseID).First(&course).Error
	if err != nil {
		return errors.New("CourseNotExisted")
	}

	result := db.Exec("INSERT IGNORE INTO bind(teacher_id,course_id) VALUES (?,?)", bind.TeacherID, bind.CourseID)
	if result.RowsAffected == 0 {
		return errors.New("CourseHasBound")
	}
	return nil
}

func UnBindCourse(unbind Bind) error {
	result := db.Delete(&unbind)
	if result.RowsAffected == 0 {
		return errors.New("CourseNotBind")
	}
	return nil
}
