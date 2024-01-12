package l18n

import "github.com/gin-gonic/gin"

var ZH_LAN = map[int]interface{}{
	ForbiddenOperate: gin.H{"error": "权限不足"},

	ParamError: gin.H{"error": "参数错误"},

	Path_ReadFail: gin.H{"error": "路径读取失败"},

	File_NotFound: gin.H{"error": "文件不存在"},

	USER_NotFound: gin.H{"error": "用户未找到"},
	User_HasExist: gin.H{"error": "用户已经存在"},

	BUCKET_NotFound: gin.H{"error": "bucket 未找到"},
	BUCKET_HasExist: gin.H{"error": "bucket 已经存在"},
}
