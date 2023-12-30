package handle

import (
	"filesystem/pojo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Auth  string `json:"auth"`
	UType int    `json:"u_type"`
}

const (
	UTypeSystem = iota + 1 // system
	UTypeAdmin             // admin
	UTypeUser              // user
)

func authJwt(ctx *gin.Context) {
	name := ctx.Request.Header.Get("user")
	auth := ctx.Request.Header.Get("auth")
	for index, user := range users {
		if user.Name == name && user.Auth == auth {
			ctx.Set("user-index", index)
			ctx.Next()
			return
		}
	}
	ctx.AbortWithStatus(401)
}

func UserRoute(r *gin.Engine) {
	userRoute := r.Group("/user", authJwt)
	{
		userRoute.POST("", CreateUser)
		userRoute.PUT("", UpdateUser)
		userRoute.DELETE("", DeleteUser)
	}
}

func getAuth(ctx *gin.Context) (int64, *User) {
	index := ctx.GetInt64("user-index")
	return index, users[index]
}

func CreateUser(ctx *gin.Context) {
	_, auth := getAuth(ctx)
	if auth.UType != UTypeSystem { // 只有system用户才有权限
		ctx.Status(403)
		return
	}
	var req pojo.CreateUser
	if errs, ok := ctx.ShouldBind(&req).(validator.ValidationErrors); ok {
		ctx.JSON(http.StatusBadRequest, errs.Translate(trans))
		return
	}
	var user = User{
		Name:  req.Name,
		Auth:  transform(randStringRunes(8)),
		UType: req.UType,
	}
	if err := data.mariadb.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	users[user.Id] = &user
	ctx.JSON(201, user)
}

func DeleteUser(ctx *gin.Context) {
	var req pojo.DeleteUser
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user User
	if err := data.mariadb.Where(&req).First(&user); err != nil {
		ctx.Status(404)
		return
	}
	_, auth := getAuth(ctx)
	if auth.UType == UTypeUser && auth.Id != user.Id || auth.UType != UTypeSystem && user.UType == UTypeSystem {
		ctx.Status(403)
		return
	}
	data.mariadb.Delete(&user)
	delete(users, user.Id)
	ctx.Status(204)
}

func UpdateUser(ctx *gin.Context) {
	var req pojo.UpdateUser
	if err := ctx.ShouldBind(&req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error_msg": fmt.Sprintf("参数错误: %v", errs.Translate(trans))})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error_msg": fmt.Sprintf("参数错误: %v", err)})
		}
		// ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user User
	data.mariadb.Where(&User{Name: req.Name}).First(&user)
	if len(req.Auth) > 0 {
		user.Auth = transform(req.Auth)
	}
	if req.UType > 0 && req.UType != user.UType {
		_, auth := getAuth(ctx)
		if req.UType < auth.UType {
			ctx.Status(403)
			return
		}
		user.UType = req.UType
	}
	data.mariadb.Where("name", req.Name).Updates(user)
	ctx.JSON(204, user)
}
