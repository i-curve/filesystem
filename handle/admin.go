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
}

func AdminRotue(r *gin.Engine) {
	r.POST("/refresh", authJwt, authSystem, func(ctx *gin.Context) {
		config.Init()
		Init()
		ctx.Status(http.StatusNoContent)
	})
}
