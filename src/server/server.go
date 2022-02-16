package server

import (
	"course_select/src/config"
	global "course_select/src/global"
	"course_select/src/rabbitmq"
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

	for i := 1; i <= 4; i++ {
		go func() {
			rabbitmq.InitConsumer()
			//TODO:
		}()
	}

	// 注册路由
	router.RegisterRouter(httpServer)

	serverError := httpServer.Run(config.GetServerConfig().HTTP_HOST + ":" + config.GetServerConfig().HTTP_PORT)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}

}
