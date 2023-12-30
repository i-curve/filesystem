package handle

import "github.com/gin-gonic/gin"

type Bucket struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	UId   int    `json:"u_id"`
	BType int    `json:"b_type"`
}

const (
	BTypeNo = iota << 0
	BTypeRead
	BTypeWrite
	BTypeReadWrite
)

func BucketRoute(r *gin.Engine) {
	bucketRoute := r.Group("/bucket", CheckBucket)
	{
		bucketRoute.POST("")
		bucketRoute.DELETE("")
	}
}
func CheckBucket(ctx *gin.Context) {}
