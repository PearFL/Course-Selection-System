package controller

import (
	global "course_select/src/global"
	"course_select/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateCourse(c *gin.Context) {

}

func GetCourse(c *gin.Context) {
	// 用于定义接受哪些请求的参数
	getCourseRequest := global.GetCourseRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&getCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.UnknownError, Message: "UnknownError"})
		return
	}

	result, err := model.GetCourse(getCourseRequest.CourseID)
	if err != nil {
		// 课程不存在
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.CourseNotExisted, Message: "CourseNotExisted"})
		return
	}

	// 成功查找到课程
	c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.OK, Data: global.TCourse{CourseID: result.CourseID, Name: result.Name,
		TeacherID: result.TeacherID}})
}

func BindCourse(c *gin.Context) {
	bindCourseRequest := global.BindCourseRequest{}

	if err := c.ShouldBind(&bindCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.UnknownError})
		return
	}

	log.Println(bindCourseRequest)

	bind := model.Bind{TeacherID: bindCourseRequest.TeacherID, CourseID: bindCourseRequest.CourseID}
	err := model.BindCourse(bind)

	if err != nil {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.CourseHasBound})
	} else {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.OK})
	}
}

func UnbindCourse(c *gin.Context) {
	unbindCourseRequest := global.UnbindCourseRequest{}

	if err := c.ShouldBind(&unbindCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.UnknownError})
		return
	}

	log.Println(unbindCourseRequest)

	unbind := model.Bind{TeacherID: unbindCourseRequest.TeacherID, CourseID: unbindCourseRequest.CourseID}
	err := model.UnBindCourse(unbind)

	if err != nil {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.CourseNotBind})
	} else {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.OK})
	}
}

func GetTeacherCourse(c *gin.Context) {

}
