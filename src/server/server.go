package server

import (
	"course_select/src/config"
	global "course_select/src/global"
	router "course_select/src/router"
	"encoding/gob"
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

	//设置session
	gob.Register(global.TMember{})
	httpServer.Use(global.GetSession())

	//建几个消息队列消费者
	go func() {
		//TODO:
	}()

	// 注册路由
	router.RegisterRouter(httpServer)

	serverError := httpServer.Run(config.GetServerConfig().HTTP_HOST + ":" + config.GetServerConfig().HTTP_PORT)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}

}
