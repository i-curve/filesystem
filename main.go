package main

import (
	"filesystem/config"
	"filesystem/handle"

	"github.com/gin-gonic/gin"
)

// go build -ldflags "-X main.Version=x.y.z"
// 2.0.0
var Version = ""

func getVersion(ctx *gin.Context) {
	ctx.String(200, "%s", Version)
}

// ApiServer api接口
func ApiServer() *gin.Engine {
	r := gin.Default()
	r.GET("/version", getVersion)
	handle.AdminRotue(r)

	handle.UserRoute(r)
	handle.BucketRoute(r)
	handle.FileRoute(r)
	return r
}

// HttpServer web 静态资源
func HttpServer() *gin.Engine {
	r := gin.Default()
	handle.StaticRoute(r)
	return r
}

func main() {
	switch config.MODE {
	case "DEBUG":
		gin.SetMode("debug")
	default:
		gin.SetMode("release")
	}

	go ApiServer().Run(":8001")
	HttpServer().Run(":8000")
}
