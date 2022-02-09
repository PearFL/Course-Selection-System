package controller

import (
	types "course_select/src/global"
	"course_select/src/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

/*
@title	CreateMember
@description	创建成员
@auth	马信宏	时间（2022/2/9   16:48 ）
*/

func CreateMember(c *gin.Context) {
	// 用于接受请求参数
	var createMemberRequest types.CreateMemberRequest

	// 接口返回的初始状态设置为OK
	var createMemberResponse types.CreateMemberResponse
	createMemberResponse.Code = types.OK

	createMemberRequest.Username = c.PostForm("Username")
	createMemberRequest.Nickname = c.PostForm("Nickname")
	createMemberRequest.Password = c.PostForm("Password")

	val, err := strconv.Atoi(c.PostForm("UserType"))

	// 枚举值(1: 管理员; 2: 学生; 3: 教师)
	if err == nil {
		if val == 1 {
			createMemberRequest.UserType = types.Admin
		} else if val == 2 {
			createMemberRequest.UserType = types.Student
		} else if val == 3 {
			createMemberRequest.UserType = types.Teacher
		} else {
			createMemberResponse.Code = types.ParamInvalid
		}
	} else {
		createMemberResponse.Code = types.ParamInvalid
	}

	// Nickname要求长度不少于8位,不超过20位，而且只支持大小写
	if len(createMemberRequest.Username) < 8 || len(createMemberRequest.Username) > 20 ||
		utils.StrIsLetter(createMemberRequest.Username) == false {
		createMemberResponse.Code = types.ParamInvalid
	}

	// Nickname要求长度不少于4位,不超过20位
	if len(createMemberRequest.Nickname) < 4 || len(createMemberRequest.Nickname) > 20 {
		createMemberResponse.Code = types.ParamInvalid
	}

	// PassWord要求长度不少于8位,不超过20位,而且同时包括大小写和数字
	if len(createMemberRequest.Password) < 8 || len(createMemberRequest.Password) > 20 ||
		utils.StrIsLowerLetterAndUpperLetterAndNumber(createMemberRequest.Password) == false {
		createMemberResponse.Code = types.ParamInvalid
	}

	// 生成uuid
	createMemberResponse.Data.UserID = uuid.NewV4().String()

	// 此处需要采用事务
	// mysql检查member表内是否该用户名已存在(Code返回2)

	// mysql写入member表内该用户

	c.JSON(http.StatusOK, gin.H{
		"Code":   createMemberResponse.Code,
		"UserID": createMemberResponse.Data.UserID,
	})

}

func GetMember(c *gin.Context) {

}

func GetMemberList(c *gin.Context) {

}

func UpdateMember(c *gin.Context) {

}

func DeleteMember(c *gin.Context) {

}
