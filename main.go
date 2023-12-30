package main

import (
	"filesystem/handle"
	"os"

	"github.com/gin-gonic/gin"
)

const version = "2.0.0"

func init() {
	handle.Init()
}

// ApiServer api接口
func ApiServer() *gin.Engine {
	r := gin.Default()
	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(200, "%s", version)
	})
	r.POST("/refresh", func(ctx *gin.Context) {
		// config.Init()
	})
	handle.UserRoute(r)
	handle.BucketRoute(r)
	handle.FileRoute(r)
	return r
}

// HttpServer web 静态资源
func HttpServer() *gin.Engine {
	r := gin.Default()
	r.Static("/", "/var/www/filesystem")
	return r
}

func main() {
	switch os.Getenv("mode") {
	case "debug":
		gin.SetMode("debug")
	default:
		gin.SetMode("release")
	}

	go ApiServer().Run(":8001")
	HttpServer().Run(":8000")
}
