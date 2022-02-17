package test

import (
	"course_select/src/database"
	"course_select/src/model"
	"course_select/src/utils"
	"fmt"
	"strconv"
)

func Test() {

	database.MySqlDb.Where("1 = 1").Delete(model.Member{})
	memberModel := model.Member{Username: "JudgeAdmin", Nickname: "JudgeAdmin",
		UserType: 1, Password: utils.Md5Encrypt("JudgePassword2022")}
	_, err := memberModel.CreateMember()

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 5000; i++ {
		memberModel := model.Member{Username: strconv.Itoa(i), Nickname: "xwytxdy",
			UserType: 2, Password: utils.Md5Encrypt("xwytxdy")}
		_, err := memberModel.CreateMember()

		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 5000; i < 6000; i++ {
		memberModel := model.Member{Username: strconv.Itoa(i), Nickname: "xwytxdy",
			UserType: 3, Password: utils.Md5Encrypt("xwytxdy")}
		_, err := memberModel.CreateMember()

		if err != nil {
			fmt.Println(err)
		}
	}

	database.MySqlDb.Where("1 = 1").Delete(model.Course{})

	for i := 0; i < 200; i++ {
		courseModel := model.Course{Name: strconv.Itoa(i), Capacity: 500}
		_, err := courseModel.CreateCourse()

		if err != nil {
			fmt.Println(err)
		}
	}
}
