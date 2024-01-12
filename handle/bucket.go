package handle

import (
	"filesystem/config"
	"filesystem/l18n"
	"filesystem/pojo"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Bucket struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	UId    int64  `json:"u_id"`
	BType  int    `json:"b_type"`
	IsTemp bool   `json:"is_temp"`
}

const (
	BTypeRead = iota + 1
	BTypeWrite
	BTypeReadWrite
)

func authBucketNo(ctx *gin.Context, bucketName string) bool {
	_, auth := getAuth(ctx)
	bucket := buckets[bucketName]
	return auth.UType == UTypeUser && bucket.UId != auth.Id
}

// func authBucket(ctx *gin.Context, bucketName string) bool {
// 	return !authBucketNo(ctx, bucketName)
// }

func StaticRoute(r *gin.Engine) {
	staticRoute := r.Group("/")
	staticRoute.Use(func(ctx *gin.Context) {
		flags := strings.Split(ctx.Request.RequestURI, "/")
		if len(flags) < 2 {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if bucket, ok := buckets[flags[1]]; !ok {
			ctx.AbortWithStatusJSON(http.StatusNotFound, lan[l18n.BUCKET_NotFound])
		} else if bucket.BType&BTypeRead == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		} else {
			ctx.Next()
		}
	})
	staticRoute.Static("", config.BASE_DIR)
}

func BucketRoute(r *gin.Engine) {
	bucketRoute := r.Group("/bucket", authJwt)
	{
		bucketRoute.GET("", listBucket)
		bucketRoute.POST("", createBucket)
		bucketRoute.DELETE("", deleteBucket)
	}
}

func listBucket(ctx *gin.Context) {
	var req pojo.ListBucket
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	_, user := getAuth(ctx)
	var buckets []*Bucket
	t := mariadb.Model(&Bucket{})
	if user.UType == UTypeUser {
		t.Where(&Bucket{UId: user.Id})
	}
	if name := strings.TrimSpace(req.Name); name != "" {
		t.Where("name like %?%", name)
	}
	t.Find(&buckets)
	ctx.JSON(http.StatusOK, buckets)
}

func createBucket(ctx *gin.Context) {
	var req pojo.CreateBucket
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if checkExistBucket(req.Name) {
		ctx.JSON(http.StatusBadRequest, lan[l18n.BUCKET_HasExist])
		return
	}
	uid, _ := getAuth(ctx)
	var bucket = Bucket{
		Name:   req.Name,
		UId:    uid,
		BType:  req.BType,
		IsTemp: req.IsTemp,
	}
	mariadb.Create(&bucket)
	buckets[req.Name] = &bucket
	mkdir(path.Join(config.BASE_DIR, req.Name))
	ctx.Status(http.StatusCreated)
}

func deleteBucket(ctx *gin.Context) {
	var req pojo.DeleteBucket
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	if !checkExistBucket(req.Name) {
		ctx.JSON(http.StatusBadRequest, lan[l18n.BUCKET_NotFound])
		return
	}
	var bucket Bucket
	mariadb.Where("name", req.Name).First(&bucket)
	_, auth := getAuth(ctx)
	if auth.UType == UTypeUser && auth.Id != bucket.UId {
		ctx.JSON(http.StatusForbidden, lan[l18n.ForbiddenOperate])
		return
	}
	mariadb.Delete(&bucket)
	delete(buckets, req.Name)
	os.RemoveAll(path.Join(config.BASE_DIR, req.Name))
	ctx.Status(http.StatusNoContent)
}

func checkExistBucket(name string) bool {
	return name != "" && mariadb.Where("name", name).First(&Bucket{}).Error == nil
}
