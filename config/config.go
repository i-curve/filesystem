package config

import (
	"fmt"
	"math/rand"
	"os"
)

type User struct {
	User string `json:"user" form:"user"`
	Auth string `json:"auth" form:"auth"`
}

var Users = make([]User, 0)
var BaseURL string = "http://127.0.0.1"

func Init() {
	baseURL := os.Getenv("BASE_URL")
	if baseURL != "" {
		BaseURL = baseURL
	}
	if os.Getenv("USER") != "" {
		Users = append(Users, User{
			User: os.Getenv("USER"),
			Auth: os.Getenv("AUTH"),
		})
	} else {
		var user = User{
			User: RandStringRunes(6),
			Auth: RandStringRunes(8),
		}
		Users = append(Users, user)
		fmt.Printf("未指定用户和认证,创建临时用户认证:\n{user:%s auth:%s}\n", user.User, user.Auth)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
