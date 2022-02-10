package controller

import (
	global "course_select/src/global"
	"course_select/src/model"
	"github.com/gin-gonic/gin"
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

}

func UnbindCourse(c *gin.Context) {

}

func GetTeacherCourse(c *gin.Context) {

}
