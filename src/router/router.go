package types

import (
	"course_select/src/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	//TEST
	g.GET("/ping", controller.Ping)

	// 成员管理
	g.POST("/member/create", controller.CreateMember)
	g.GET("/member", controller.GetMember)
	g.GET("/member/list", controller.GetMemberList)
	g.POST("/member/update", controller.UpdateMember)
	g.POST("/member/delete", controller.DeleteMember)

	// 登录

	g.POST("/auth/login")
	g.POST("/auth/logout")
	g.GET("/auth/whoami")

	// 排课
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")

}
