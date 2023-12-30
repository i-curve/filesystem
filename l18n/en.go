package l18n

import "github.com/gin-gonic/gin"

const (
	ForbiddenOperate = iota + 1
	USER_NotFound
	User_HasExist

	BUCKET_NotFound
	BUCKET_HasExist
)

var EN_LAN = map[int]interface{}{
	ForbiddenOperate: gin.H{"error": "operator forbidden"},

	USER_NotFound: gin.H{"error": "user not found"},
	User_HasExist: gin.H{"error": "user has exist"},

	BUCKET_NotFound: gin.H{"error": "bucket not found"},
	BUCKET_HasExist: gin.H{"error": "bucket has exist"},
}
