package handle

import (
	"filesystem/config"
	"filesystem/l18n"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type fileUpload struct {
	Bucket   string                `json:"bucket" form:"bucket" binding:"required"`
	PATH     string                `json:"path" form:"path" binding:"required"`
	File     *multipart.FileHeader `json:"file" form:"file" binding:"required"`
	Duration int64                 `json:"duration" form:"duration"`
}

type fileDelete struct {
	Bucket string `json:"bucket" form:"bucket" binding:"required"`
	PATH   string `json:"path" form:"path" binding:"required"`
}

type fileDownload struct {
	Bucket string `json:"bucket" form:"bucket" binding:"required"`
	PATH   string `json:"path" form:"path" binding:"required"`
}

type fileCopy struct {
	SBucket string `json:"s_bucket" form:"s_bucket" binding:"required"`
	SPath   string `json:"s_path" form:"s_path" binding:"required"`
	DBucket string
	DPath   string `json:"d_path" form:"d_path" binding:"required"`
}

func FileRoute(r *gin.Engine) {
	fileRoute := r.Group("file", authJwt)
	{
		contact := Filesystem{}
		fileRoute.POST("/create", contact.create)    // 文件上传
		fileRoute.GET("/download", contact.download) // 文件下载
		fileRoute.POST("/copy", contact.copy)        // 文件复制
		fileRoute.POST("/move", contact.move)        // 文件转移
		fileRoute.DELETE("/delete", contact.delete)  // 文件删除
	}
}

type Filesystem struct{}

// 文件上传
func (f Filesystem) create(ctx *gin.Context) {
	var req fileUpload
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistBucket(req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.BUCKET_NotFound])
		return
	}

	if authBucketNo(ctx, req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}
	r, _ := req.File.Open()
	writeFile(req.Bucket, req.PATH, r)
	if buckets[req.Bucket].IsTemp {
		req.Duration = config.TEMP_DURATION
	}
	if req.Duration > 0 {
		var cron = Cron{
			Bucket:     req.Bucket,
			Path:       req.PATH,
			DeleteTime: time.Now().Add(time.Second * time.Duration(req.Duration)),
		}
		mariadb.Create(&cron)
		CronDelete(&cron)
	}
}

// 文件获取
func (f Filesystem) download(ctx *gin.Context) {
	var req fileDownload
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistBucket(req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.BUCKET_NotFound])
		return
	}
	if authBucketNo(ctx, req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}

	file, err := os.Open(path.Join(config.BASE_DIR, req.Bucket, req.PATH))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+path.Base(req.PATH))
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	io.Copy(ctx.Writer, file)
}

// 文件转移
func (f Filesystem) move(ctx *gin.Context) {
	var req fileCopy
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistFile(req.SBucket, req.SPath) {
		ctx.JSON(http.StatusBadRequest, lan[l18n.File_NotFound])
		return
	}
	if authBucketNo(ctx, req.SBucket) || authBucketNo(ctx, req.DBucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}
	moveFile(req.SBucket, req.SPath, req.DBucket, req.DPath)
	ctx.Status(http.StatusOK)
}

// 文件复制
func (f Filesystem) copy(ctx *gin.Context) {
	var req fileCopy
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistFile(req.SBucket, req.SPath) {
		ctx.JSON(http.StatusBadRequest, lan[l18n.File_NotFound])
		return
	}
	if authBucketNo(ctx, req.SBucket) || authBucketNo(ctx, req.DBucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}
	copyFile(req.SBucket, req.SPath, req.DBucket, req.DPath)
	ctx.Status(http.StatusOK)
}

// 文件删除
func (f Filesystem) delete(ctx *gin.Context) {
	var req fileDelete
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistBucket(req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.BUCKET_NotFound])
		return
	}
	if authBucketNo(ctx, req.Bucket) {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}
	removeFile(req.Bucket, req.PATH)
	ctx.Status(http.StatusNoContent)
}
