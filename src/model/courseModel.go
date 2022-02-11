package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	CourseID    string `json:"course_id" form:"course_id" gorm:"primary_key"`
	TeacherID   string `json:"teacher_id" form:"teacher_id"`
	Name        string `json:"name" form:"name"`
	Capacity    int    `json:"capacity" form:"capacity"`
	CapSelected int    `json:"cap_select" form:"cap_select"`
}

func (Course) TableName() string {
	return "course"
}

func (course *Course) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("course_id", uuid)
}

func (course *Course) CreateCourse() (string, error) {
	err := db.Create(&course).Error
	if err != nil {
		return "", err
	}
	return course.CourseID, nil
}

func (course *Course) GetCourse(id string) (Course, error) {
	var ans Course

	err := db.Model(&Course{}).Where("course_id = ?", id).First(&ans).Error

	if err != nil {
		return ans, err
	}

	return ans, nil
}

// GetCourses 得到指定老师的课程
func (course *Course) GetCourses(tid string) ([]Course, error) {
	var bindCourseOfTeacher []Bind
	var ans []Course
	//得到该老师绑定的课程ID
	err := db.Where("teacher_id = ?", tid).Find(&bindCourseOfTeacher).Error
	if err != nil {
		return ans, err
	}
	//取出每个课程ID对应的Course记录
	for _, v := range bindCourseOfTeacher {
		var temp []Course
		err := db.Where("course_id = ?", v.CourseID).Find(&temp).Error
		if err != nil {
			return ans, err
		}
		ans = append(ans, temp...)
	}

	return ans, nil
}
