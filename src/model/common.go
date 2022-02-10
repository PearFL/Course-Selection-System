package model

import (
	"course_select/src/database"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = database.MySqlDb
