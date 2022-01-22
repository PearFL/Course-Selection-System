package server

import (
	"course_select/src/config"
	router "course_select/src/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine) {

	// 生成日志
	logFile, _ := os.Create(config.GetLogPath())
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout, os.Stdin, os.Stderr)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(config.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 注册路由
	router.RegisterRouter(httpServer)

	serverError := httpServer.Run(config.GetServerConfig().HTTP_HOST + ":" + config.GetServerConfig().HTTP_PORT)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}

}
