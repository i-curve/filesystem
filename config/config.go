package config

import (
	"fmt"
	"os"
	"strings"
)

var internal_env = map[string]string{
	"LANGUAGE":       "zh", // en|zh
	"MODE":           "RELEASE",
	"BASE_DIR":       "/var/www/filesystem",
	"DSN":            "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	"MYSQL_HOST":     "127.0.0.1",
	"MYSQL_USER":     "root",
	"MYSQL_PASSWORD": "123456",
	"MYSQL_PORT":     "3306",
	"DATABASE":       "filesystem",
}
var MODE = ""
var BASE_DIR = ""
var DSN = ""
var LANGUAGE = ""

func init() {
	Init()
}

func Init() {
	for _, env := range os.Environ() {
		items := strings.Split(env, "=")
		if _, ok := internal_env[items[0]]; ok {
			internal_env[items[0]] = items[1]
		}
	}
	MODE = internal_env["MODE"]
	LANGUAGE = internal_env["LANGUAGE"]
	BASE_DIR = internal_env["BASE_DIR"]
	DSN = fmt.Sprintf(internal_env["DSN"], internal_env["MYSQL_USER"], internal_env["MYSQL_PASSWORD"],
		internal_env["MYSQL_HOST"], internal_env["MYSQL_PORT"], internal_env["DATABASE"])
}
