package main

import (
	"course_select/src/database"
	global "course_select/src/global"
	"course_select/src/rabbitmq"
	"course_select/src/server"
	"github.com/gin-gonic/gin"
)

func main() {

	defer func() {
		database.MySqlDb.Close()
		database.RedisClient.Close()
		rabbitmq.RabbitMQConn.Close()
		rabbitmq.RabbitMQChannel.Close()
		global.LogFile.Close()
	}()

	httpServer := gin.Default()
	server.Run(httpServer)

	// test.Test()

}
