package handle

import (
	"filesystem/l18n"
	"filesystem/pojo"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Bucket struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	UId   int64  `json:"u_id"`
	BType int    `json:"b_type"`
}

const (
	BTypeRead = iota + 1
	BTypeWrite
	BTypeReadWrite
)

func BucketRoute(r *gin.Engine) {
	bucketRoute := r.Group("/bucket", authJwt)
	{
		bucketRoute.POST("", createBucket)
		bucketRoute.DELETE("", deleteBucket)
	}
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
	mariadb.Create(&Bucket{
		Name:  req.Name,
		UId:   uid,
		BType: req.BType,
	})
	os.MkdirAll(path.Join(BASE_DIR, req.Name), os.ModePerm)
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
	os.RemoveAll(path.Join(BASE_DIR, req.Name))
	ctx.Status(http.StatusNoContent)
}

func checkExistBucket(name string) bool {
	if name == "" {
		return false
	}
	return mariadb.Where(Bucket{
		Name: name,
	}).First(&Bucket{}).Error == nil
}
