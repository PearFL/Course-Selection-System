package validate

import (
	global "course_select/src/global"
)

var CourseValidate global.Validator

func init() {
	rules := map[string]string{
		"Name":     "required",
		"Cap":      "required|int:1,",
		"CourseID": "required",
	}

	scenes := map[string][]string{
		"add": {"Name", "Cap"},
		"get": {"CourseID"},
	}

	CourseValidate.Rules = rules
	CourseValidate.Scenes = scenes
}
