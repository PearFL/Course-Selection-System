package controller

import (
	"course_select/src/database"
	global "course_select/src/global"
	"course_select/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

var mutex = sync.Mutex{}

// var limiter = rate.NewLimiter(1000, 3000)

func BookCourse(c *gin.Context) {
	// 十秒后再请求
	// c.Set("Deadline", time.Now().Add(time.Second*10))
	// limiter.Wait(c)

	rc := database.RedisClient.Get()
	defer rc.Close()

	// 用于定义接受哪些请求的参数
	bookCourseRequest := global.BookCourseRequest{}

	// 用于定义获取参数值
	if err := c.ShouldBind(&bookCourseRequest); err != nil {
		c.JSON(http.StatusOK, global.BookCourseResponse{Code: global.UnknownError})
		return
	}

	// 校验这是不是学生
	if !model.IsStudentLegal(bookCourseRequest.StudentID, rc) {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.StudentNotExisted})
		return
	}

	// 校验课程存不存在
	if !model.IsCourseLegal(bookCourseRequest.CourseID, rc) {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.CourseNotExisted})
		return
	}

	// 校验学生是否选过这个课
	if model.IsBooked(bookCourseRequest.StudentID, bookCourseRequest.CourseID, rc) {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.RepeatRequest})
		return
	}

	// 秒杀减库存
	cnt := model.DecrAndGet(bookCourseRequest.CourseID, rc)
	if cnt < 0 {
		// 超卖直接回滚
		model.IncrAndGet(bookCourseRequest.CourseID, rc)
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.CourseNotAvailable})
		return
	}

	// redis写库
	model.UpdateStudentCourse(bookCourseRequest.StudentID, bookCourseRequest.CourseID, rc)

	go func() {
		mutex.Lock()
		err := InitProducer(global.BookCourseRequest{
			StudentID: bookCourseRequest.StudentID,
			CourseID:  bookCourseRequest.CourseID,
		})
		mutex.Unlock()
		if err != nil {
			log.Println("消息队列错误")
			return
		}
	}()

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

	log.Println(studentCourseRequest)

	// 校验这是不是学生
	if !model.IsStudentLegal(studentCourseRequest.StudentID, rc) {
		c.JSON(http.StatusOK, global.ResponseMeta{Code: global.StudentNotExisted})
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
