package controller

import (
	global "course_select/src/global"
	"course_select/src/model"
	"course_select/src/validate"
	"github.com/gin-gonic/gin"
	"log"
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
		c.JSON(http.StatusOK, global.UnbindCourseResponse{Code: global.UnknownError})
		return
	}

	log.Println(unbindCourseRequest)

	unbind := model.Bind{TeacherID: unbindCourseRequest.TeacherID, CourseID: unbindCourseRequest.CourseID}
	err := model.UnBindCourse(unbind)

	if err != nil {
		c.JSON(http.StatusOK, global.UnbindCourseResponse{Code: global.CourseNotBind})
	} else {
		c.JSON(http.StatusOK, global.UnbindCourseResponse{Code: global.OK})
	}
}

func GetTeacherCourse(c *gin.Context) {

}
