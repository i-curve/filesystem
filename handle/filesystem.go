package handle

import (
	"filesystem/config"
	"filesystem/model"
	"filesystem/util"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// 文件上传
func UploadFile(ctx *gin.Context) {
	var request model.Request
	ctx.ShouldBind(&request)
	if data, err := util.WriteFile(request.Path, request.File); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": err.HTTPCode(), "err": err.Error(), "msg": err.Message()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 201, "url": data.URL, "short_url": data.ShortURL})
	}
}

// 文件获取
func GetFile(ctx *gin.Context) {
	var request model.Request
	ctx.ShouldBind(&request)
	var base string
	if request.ShortURL != "" {
		base = request.ShortURL
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "url": path.Join(config.BaseURL, base)})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "err": "路径不能为空", "msg": "路径不能为空"})
	}
}

// 文件转移
func MoveFile(ctx *gin.Context) {
	var request model.Request
	ctx.ShouldBind(&request)
	if err := util.MoveFile(request.ShortURL, request.NewURL); err != nil {
		ctx.JSON(err.HTTPCode(), gin.H{"code": err.Code(), "err": err.Error(), "msg": err.Message()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 201})
	}
}

// 文件复制
func CopyFile(ctx *gin.Context) {
	var request model.Request
	ctx.ShouldBind(&request)
	if err := util.CopyFile(request.ShortURL, request.NewURL); err != nil {
		ctx.JSON(err.HTTPCode(), gin.H{"code": err.Code(), "err": err.Error(), "msg": err.Message()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 201})
	}
}

// 文件删除
func DeleteFile(ctx *gin.Context) {
	var request model.Request
	ctx.ShouldBind(&request)
	if err := util.DeleteFile(request.ShortURL); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": err.HTTPCode(), "err": err.Error(), "msg": err.Message()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 204})
	}
}
