package controller

import (
	global "course_select/src/global"
	"course_select/src/model"
	"course_select/src/validate"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	getTeacherCourseRequest := global.GetTeacherCourseRequest{}
	courseModel := model.Course{}
	if err := c.ShouldBind(&getTeacherCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.GetTeacherCourseResponse{Code: global.UnknownError})
		return
	}
	courses, err := courseModel.GetCourses(getTeacherCourseRequest.TeacherID)
	if err != nil {
		c.JSON(http.StatusOK, global.GetTeacherCourseResponse{Code: global.UnknownError})
		return
	}
	CourseList := make([]*global.TCourse, len(courses))
	for i, v := range courses {
		CourseList[i] = &global.TCourse{
			CourseID:  v.CourseID,
			Name:      v.Name,
			TeacherID: v.TeacherID,
		}
	}
	c.JSON(http.StatusOK, global.GetTeacherCourseResponse{
		Code: global.OK,
		Data: struct{ CourseList []*global.TCourse }{CourseList: CourseList}})
}

func ScheduleCourse(c *gin.Context) {
	scheduleCourseRequest := global.ScheduleCourseRequest{}
	if err := c.ShouldBind(&scheduleCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.ScheduleCourseResponse{Code: global.UnknownError})
		return
	}

	g := scheduleCourseRequest.TeacherCourseRelationShip

	match := make(map[string]string, len(g))
	ans := make(map[string]string, len(g))

	cnt := 0
	for i := range match {
		match[i] = ""
	}
	var used map[string]bool
	var f func(string) bool
	f = func(v string) bool {
		used[v] = true
		for _, w := range g[v] {
			if mw := match[w]; mw == "" || !used[mw] && f(mw) {
				match[w] = v
				match[v] = w

				ans[v] = w

				return true
			}
		}
		return false
	}

	for v := range g {
		if match[v] == "" {
			used = make(map[string]bool, len(g))
			if f(v) {
				cnt++ // +=2
			}
		}
	}
	c.JSON(http.StatusOK, global.ScheduleCourseResponse{Code: global.OK, Data: ans})
}
