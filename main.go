package main

import (
	"filesystem/config"
	"filesystem/handle"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 版本信息
var version = "1.0.0"

func init() {
	config.Init()
}

func main() {
	gin.SetMode("release")
	router := gin.Default()

	// 获取版本信息
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"version": version})
	})
	// 重新加载配置文件
	router.POST("/reload", func(ctx *gin.Context) {
		config.Init()
	})

	// 注册回调
	routes := router.Group("", handle.Middleware)
	{
		routes.POST("/upload", handle.UploadFile) // 文件上传
		routes.GET("/file", handle.GetFile)       // 获取文件url
		routes.POST("/copy", handle.CopyFile)     // 文件复制
		routes.POST("/move", handle.MoveFile)     // 文件转移
		routes.DELETE("/file", handle.DeleteFile) // 文件删除
	}

	// 启动项目
	router.Run()
}
