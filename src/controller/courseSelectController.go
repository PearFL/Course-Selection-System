package controller

import (
	"course_select/src/database"
	global "course_select/src/global"
	"course_select/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookCourse(c *gin.Context) {
	rc := database.RedisClient.Get()
	defer rc.Close()

	// 用于定义接受哪些请求的参数
	bookCourseRequest := global.BookCourseRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&bookCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.UnknownError})
		return
	}

	// 校验学生是否选过这个课
	if model.IsBooked(bookCourseRequest.StudentID, bookCourseRequest.CourseID, rc) {
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.StudentHasNoCourse})
		return
	}

	// 秒杀减库存
	cnt := model.DecrAndGet(bookCourseRequest.CourseID, rc)
	if cnt < 0 {
		// 超卖直接回滚
		model.IncrAndGet(bookCourseRequest.CourseID, rc)
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.CourseNotAvailable})
		return
	}

	// redis写库
	model.UpdateStudentCourse(bookCourseRequest.StudentID, bookCourseRequest.CourseID, rc)

	// 生产者生产消息
	err := InitProducer(global.BookCourseRequest{
		StudentID: bookCourseRequest.StudentID,
		CourseID:  bookCourseRequest.CourseID,
	})
	if err != nil {
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.UnknownError})
		return
	}

	c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.OK})
	return
}

func GetStudentCourse(c *gin.Context) {
	rc := database.RedisClient.Get()
	defer rc.Close()

	// 用于定义接受哪些请求的参数
	studentCourseRequest := global.GetStudentCourseRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&studentCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.GetStudentCourseResponse{Code: global.UnknownError})
		return
	}

	strings := model.GetStudentCourses(studentCourseRequest.StudentID, rc)

	var courses = make([]global.TCourse, len(strings))

	// 从redis中提取课程信息组成response中的Data
	for i, v := range strings {
		courses[i] = global.TCourse{
			CourseID:  v,
			Name:      model.GetCourseNameById(v, rc),
			TeacherID: model.GetTeacherByCourseId(v, rc),
		}
	}

	c.JSON(http.StatusOK, global.GetStudentCourseResponse{Code: global.OK, Data: struct{ CourseList []global.TCourse }{CourseList: courses}})
}
