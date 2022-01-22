package main

import (
	"course_select/src/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello World")
	httpServer := gin.Default()
	server.Run(httpServer)
}
