package controller

import (
	"course_select/src/database"
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

	result := database.MySqlDb.Exec("INSERT IGNORE INTO bind(teacher_id,course_id) VALUES (?,?)", bindCourseRequest.TeacherID, bindCourseRequest.CourseID)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, global.BindCourseResponse{Code: global.OK})
	} else {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.CourseHasBound, Message: "CourseHasBound"})
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
	result := database.MySqlDb.Delete(&unbind)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, global.UnbindCourseResponse{Code: global.OK})
	} else {
		c.JSON(http.StatusOK, global.ErrorResponse{Code: global.CourseNotBind, Message: "CourseNotBind"})
	}
}

func GetTeacherCourse(c *gin.Context) {

}
