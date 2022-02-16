package validate

import (
	global "course_select/src/global"
)

var CourseValidate global.Validator

func init() {
	rules := map[string]string{
		"Name":      "required",
		"Cap":       "required|int:1,",
		"CourseID":  "required",
		"TeacherID": "required",
	}

	scenes := map[string][]string{
		"add":        {"Name", "Cap"},
		"get":        {"CourseID"},
		"bind":       {"CourseID", "TeacherID"},
		"unbind":     {"CourseID", "TeacherID"},
		"get_course": {"TeacherID"},
	}

	CourseValidate.Rules = rules
	CourseValidate.Scenes = scenes
}
