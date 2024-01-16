package handle

import (
	"filesystem/config"
	"filesystem/l18n"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	Init()
}

func Init() {
	lan = l18n.ZH_LAN
	initDir()
	initTrans()
	initDB()
	initData()
	initCron()
}

func authJwtOrNot(ctx *gin.Context) {
	name := ctx.Request.Header.Get("user")
	auth := ctx.Request.Header.Get("auth")
	if name == "" && auth == "" {
		ctx.Next()
		return
	}
	for index, user := range users {
		if user.Name == name && user.Auth == auth {
			ctx.Set("user-index", index)
			ctx.Next()
			return
		}
	}
	ctx.AbortWithStatus(401)
}

func AdminRotue(r *gin.Engine, version string) {
	r.POST("/refresh", authJwt, authSystem, func(ctx *gin.Context) {
		config.Init()
		Init()
		ctx.Status(http.StatusNoContent)
	})
	r.GET("/version", authJwtOrNot, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, version)
	})
}
