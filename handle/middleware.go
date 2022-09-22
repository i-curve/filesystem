package handle

import (
	"filesystem/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	var user config.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 401, "err": "用户未授权", "msg": "用户未授权"})
		return
	}

	var flag bool
	for _, v := range config.Users {
		if v.User == user.User && v.Auth == user.Auth {
			flag = true
			break
		}
	}

	if flag {
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 401, "err": "用户未授权", "msg": "用户未授权"})
		c.Abort()
	}
}
