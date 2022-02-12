package types

import (
	"course_select/src/controller"
	global "course_select/src/global"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	//TEST
	// g.GET("/ping", controller.Ping)

	//需要鉴权的路由组
	authGroup := g.Group("")
	authGroup.Use(global.BackendAuth())
	{
		authGroup.POST("/member/create", controller.CreateMember)
		authGroup.POST("/course/create", controller.CreateCourse)
	}

	// 成员管理
	// g.POST("/member/create", controller.CreateMember)
	g.GET("/member", controller.GetMember)
	g.GET("/member/list", controller.GetMemberList)
	g.POST("/member/update", controller.UpdateMember)
	g.POST("/member/delete", controller.DeleteMember)

	// 登录
	g.POST("/auth/login", controller.Login)
	g.POST("/auth/logout", controller.Logout)
	g.GET("/auth/whoami", controller.WhoAmI)

	// 排课
	// g.POST("/course/create", controller.CreateCourse)
	g.GET("/course/get", controller.GetCourse)

	g.POST("/teacher/bind_course", controller.BindCourse)
	g.POST("/teacher/unbind_course", controller.UnbindCourse)
	g.GET("/teacher/get_course", controller.GetTeacherCourse)
	g.POST("/course/schedule", controller.ScheduleCourse)

	// 抢课
	g.POST("/student/book_course", controller.BookCourse)
	g.GET("/student/course", controller.GetStudentCourse)

}
