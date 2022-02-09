package controller

import (
	global "course_select/src/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@title	CreateMember
@description	创建成员
@auth	马信宏	时间（2022/2/9   16:48 ）
*/

func CreateMember(c *gin.Context) {
	// 用于接受请求参数
	createMemberRequest := global.CreateMemberRequest{}

	// memberModel := model.Member{}

	if err := c.ShouldBind(&createMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.UnknownError})
		return
	}

	//TODO:这里用中间件检验参数是否符合要求

	//FIXME:我不大看得懂这几行里是什么意思，如果是鉴权的话交给路由中间件
	// val, err := strconv.Atoi(c.PostForm("UserType"))

	// // 枚举值(1: 管理员; 2: 学生; 3: 教师)
	// if err == nil {
	// 	if val == 1 {
	// 		createMemberRequest.UserType = types.Admin
	// 	} else if val == 2 {
	// 		createMemberRequest.UserType = types.Student
	// 	} else if val == 3 {
	// 		createMemberRequest.UserType = types.Teacher
	// 	} else {
	// 		createMemberResponse.Code = types.ParamInvalid
	// 	}
	// } else {
	// 	createMemberResponse.Code = types.ParamInvalid
	// }

	// TODO:sql,此处需要采用事务
	// mysql检查member表内是否该用户名已存在(Code返回2)  mysql写入member表内该用户
	// userID, err := memberModel.CreateMember(/*参数*/)

	// if err != nil {
	// 	c.JSON(http.StatusOK, global.CreateMemberResponse{Code: 255 /*这里具体看是学生已存在还是怎么个其他错误*/})
	// }

	c.JSON(http.StatusOK, global.CreateMemberResponse{Code: 0, Data: struct{ UserID string }{ /*数据库返回参数*/ }})

}

func GetMember(c *gin.Context) {

}

func GetMemberList(c *gin.Context) {

}

func UpdateMember(c *gin.Context) {

}

func DeleteMember(c *gin.Context) {

}
