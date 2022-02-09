package validate

import (
	global "course_select/src/global"
)

var MemberValidate global.Validator

func init() {
	rules := map[string]string{
		"Username": "required|minLen:4|maxLen:20|alpha",
		"Nickname": "required|minLen:8|maxLen:20",
		"UserType": "required",
		"Password": "required|alphaNum|string:8,20|passwordValidator",
	}

	scenes := map[string][]string{}

	MemberValidate.Rules = rules
	MemberValidate.Scenes = scenes
}
