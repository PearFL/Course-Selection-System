package model

import (
	"course_select/src/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = database.MySqlDb

func SaveChoice(studentID, courseID string) error {

	return db.Transaction(func(tx *gorm.DB) error {
		var course Course
		if err := tx.Where("course_id = ?", courseID).First(&course).Error; err != nil {
			return err
		}

		if err := db.Model(&Course{}).Where("course_id = ?", courseID).Update(Course{CourseID: course.CourseID, Capacity: course.Capacity, CapSelected: course.CapSelected + 1, Name: course.Name}).Error; err != nil {
			return err
		}

		if err := tx.Create(&Choice{StudentID: studentID, CourseID: courseID}).Error; err != nil {
			return err
		}

		return nil
	})
}
