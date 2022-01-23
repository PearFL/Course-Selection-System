package main

import (
	"course_select/src/database"
	"course_select/src/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello World")

	defer func() {
		database.MySqlDb.Close()
		database.RedisClient.Close()
	}()

	httpServer := gin.Default()
	server.Run(httpServer)

}
