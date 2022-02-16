package types

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cookiesName string = "camp-session"

func BackendAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println(c.FullPath())

		session := sessions.Default(c)
		sessionId, err := c.Cookie(cookiesName)

		if err != nil {
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}

		v := session.Get(sessionId)
		if v == nil {
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}
		// log.Println(v)
		user := v.(TMember)

		if user.UserType != 1 {
			c.JSON(http.StatusOK, ResponseMeta{Code: PermDenied})
			c.Abort()
			return
		}
		c.Next()
	}
}
