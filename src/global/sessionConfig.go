package types

import (
	"course_select/src/database"
	"github.com/gin-contrib/sessions"
	redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func GetSession() gin.HandlerFunc {
	store, _ := redis.NewStoreWithPool(database.RedisClient, []byte("passord"))
	return sessions.Sessions("MySession", store)
}
