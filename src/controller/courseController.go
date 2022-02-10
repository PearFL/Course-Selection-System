package controller

import (
	"course_select/src/database"
	global "course_select/src/global"
	"course_select/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"course_select/src/validate"
	"net/http"
)

func CreateCourse(c *gin.Context) {
	request := global.CreateCourseRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.UnknownError})
		return
	}

	requestMap := global.Struct2Map(request)
	courseValidate := validate.CourseValidate
	res, _ := courseValidate.ValidateMap(requestMap, "add")
	if !res {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.ParamInvalid})
		return
	}

	courseModel := model.Course{Name: request.Name, Capacity: request.Cap}
	uuid, err := courseModel.CreateCourse()
	if err != nil {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.UnknownError})
		return
	}
	c.JSON(http.StatusOK, global.CreateCourseResponse{Code: global.OK, Data: struct{ CourseID string }{uuid}})
}

func GetCourse(c *gin.Context) {
	request := global.GetCourseRequest{}
	courseModel := model.Course{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.UnknownError})
		return
	}
	requestMap := global.Struct2Map(request)
	courseValidate := validate.CourseValidate
	res, _ := courseValidate.ValidateMap(requestMap, "get")
	if !res {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.ParamInvalid})
		return
	}
	course, err := courseModel.GetCourse(request.CourseID)
	if err != nil {
		c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.CourseNotExisted})
	}
	c.JSON(http.StatusOK, global.GetCourseResponse{Code: global.OK, Data: global.TCourse{CourseID: course.CourseID, Name: course.Name, TeacherID: course.TeacherID}})

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
