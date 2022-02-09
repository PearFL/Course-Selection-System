package types

import "github.com/gin-gonic/gin"

//TODO:晚上看一下要不要使用
func ReturnType(code ErrNo, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"data": data,
	}
}
