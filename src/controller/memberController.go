package controller

import (
	global "course_select/src/global"
	"course_select/src/model"
	"course_select/src/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	createMemberRequest := global.CreateMemberRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&createMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.UnknownError})
		return
	}

	requestMap := global.Struct2Map(createMemberRequest)
	memberValidate := validate.MemberValidate
	res, _ := memberValidate.ValidateMap(requestMap, "add")

	if !res {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.ParamInvalid, Message: "ParamInvalid"})
		return
	}

	if createMemberRequest.PasswordValidator(createMemberRequest.Password) == false {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.ParamInvalid, Message: "ParamInvalid"})
		return
	}

	if createMemberRequest.UserTypeValidator(createMemberRequest.UserType) == false {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.ParamInvalid, Message: "ParamInvalid"})
		return
	}

	memberModel := model.Member{Username: createMemberRequest.Username, Nickname: createMemberRequest.Nickname,
		UserType: createMemberRequest.UserType, Password: createMemberRequest.Password}
	uuid, err := memberModel.CreateMember()

	if err != nil {
		if err.Error() == "UserHasExisted" {
			c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.UserHasExisted})
		} else {
			c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.UnknownError})
		}
		return
	}

	c.JSON(http.StatusOK, global.CreateMemberResponse{Code: global.OK, Data: struct{ UserID string }{uuid}})

}

func GetMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	getMemberRequest := global.GetMemberRequest{}
	memberModel := model.Member{}
	// 用于定义获取参数值
	if err := c.ShouldBind(&getMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UnknownError, Message: "UnknownError"})
		return
	}

	requestMap := global.Struct2Map(getMemberRequest)

	fmt.Println(requestMap)

	log.Println(getMemberRequest)

	result, err := memberModel.GetMember(getMemberRequest.UserID)
	if err != nil {
		// 用户不存在
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UserNotExisted, Message: "UserNotExisted"})
		return
	}

	if result.IsDeleted == true {
		// 用户已经被删除
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UserHasDeleted, Message: "UserHasDeleted"})
		return
	}

	// 成功查找到用户
	c.JSON(http.StatusOK, global.GetMemberResponse{Code: global.OK, Data: global.TMember{UserID: result.UserID, Nickname: result.Nickname,
		Username: result.Username, UserType: result.UserType}})
}

func GetMemberList(c *gin.Context) {
	// 获取参数
	GetMemberListRequest := global.GetMemberListRequest{}
	memberModel := model.Member{}
	if err := c.ShouldBind(&GetMemberListRequest); err != nil {
		c.JSON(http.StatusOK, global.GetMemberListResponse{Code: global.UnknownError})
		return
	}

	offset, limit := GetMemberListRequest.Offset, GetMemberListRequest.Limit

	// 查询并取出结果
	members, err := memberModel.GetAllMembers(offset, limit)
	if err != nil {
		c.JSON(http.StatusOK, global.GetMemberListResponse{Code: global.UnknownError})
		return
	}

	MemberList := make([]global.TMember, len(members))
	for i, v := range members {
		MemberList[i] = global.TMember{
			UserID:   v.UserID,
			Nickname: v.Nickname,
			Username: v.Username,
			UserType: v.UserType,
		}
	}
	// 返回
	c.JSON(http.StatusOK, global.GetMemberListResponse{
		Code: global.OK,
		Data: struct{ MemberList []global.TMember }{MemberList: MemberList}})
}

func UpdateMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	updateMemberRequest := global.UpdateMemberRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&updateMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UnknownError, Message: "UnknownError"})
		return
	}

	log.Println(updateMemberRequest)

	err := model.UpdateMember(updateMemberRequest.UserID, updateMemberRequest.Nickname)

	if err != nil {
		// 用户不存在
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UserNotExisted, Message: "UserNotExisted"})
		return
	}

	// 成功更新用户昵称
	c.JSON(http.StatusOK, global.UpdateMemberResponse{Code: global.OK})
}

func DeleteMember(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	deleteMemberRequest := global.DeleteMemberRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&deleteMemberRequest); err != nil {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UnknownError, Message: "UnknownError"})
		return
	}

	log.Println(deleteMemberRequest)

	err := model.DeleteMember(deleteMemberRequest.UserID)

	if err != nil {
		// 用户不存在
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UserNotExisted, Message: "UserNotExisted"})
		return
	}

	// 成功删除用户
	c.JSON(http.StatusOK, global.DeleteMemberResponse{Code: global.OK})
}
