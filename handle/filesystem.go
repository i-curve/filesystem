package handle

import (
	"filesystem/config"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type Request struct {
	URL      string                `json:"url" form:"url"`
	ShortURL string                `json:"short_url" form:"short_url"`
	NewURL   string                `json:"new_url" form:"new_url"`
	Path     string                `json:"path" form:"path"`
	File     *multipart.FileHeader `json:"file" form:"file"`
}
type Reply struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}

func FileRoute(r *gin.Engine) {
	fileRoute := r.Group("", Middleware)
	{
		fileRoute.POST("/upload", UploadFile) // 文件上传
		fileRoute.GET("/file", GetFile)       // 获取文件
		// fileRoute.POST("/copy", CopyFile)     // 文件复制
		// fileRoute.POST("/move", MoveFile) // 文件转移
		// fileRoute.DELETE("/file", DeleteFile) // 文件删除
	}
}

// 文件上传
func UploadFile(ctx *gin.Context) {
	var request Request
	ctx.ShouldBind(&request)
	// if data, err := util.WriteFile(request.Path, request.File); err != nil {
	// 	ctx.JSON(http.StatusOK, gin.H{"code": err.HTTPCode(), "err": err.Error(), "msg": err.Message()})
	// } else {
	// 	ctx.JSON(http.StatusOK, gin.H{"code": 201, "url": data.URL, "short_url": data.ShortURL})
	// }
}

// 文件获取
func GetFile(ctx *gin.Context) {
	var request Request
	ctx.ShouldBind(&request)
	var base string
	if request.ShortURL != "" {
		base = request.ShortURL
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "url": path.Join(config.BaseURL, base)})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "err": "路径不能为空", "msg": "路径不能为空"})
	}
}

// // 文件转移
// func MoveFile(ctx *gin.Context) {
// 	var request Request
// 	ctx.ShouldBind(&request)
// 	if err := moveFile(request.ShortURL, request.NewURL); err != nil {
// 		ctx.JSON(err.HTTPCode(), gin.H{"code": err.Code(), "err": err.Error(), "msg": err.Message()})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"code": 201})
// 	}
// }

// // 文件复制
// func CopyFile(ctx *gin.Context) {
// 	var request Request
// 	ctx.ShouldBind(&request)
// 	if err := util.CopyFile(request.ShortURL, request.NewURL); err != nil {
// 		ctx.JSON(err.HTTPCode(), gin.H{"code": err.Code(), "err": err.Error(), "msg": err.Message()})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"code": 201})
// 	}
// }

// // 文件删除
// func DeleteFile(ctx *gin.Context) {
// 	var request Request
// 	ctx.ShouldBind(&request)
// 	if err := util.DeleteFile(request.ShortURL); err != nil {
// 		ctx.JSON(http.StatusOK, gin.H{"code": err.HTTPCode(), "err": err.Error(), "msg": err.Message()})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"code": 204})
// 	}
// }
