package handle

import (
	"crypto/md5"
	"encoding/hex"
	"filesystem/l18n"
	"fmt"
	"math/rand"
	"os"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var trans ut.Translator
var BASE_DIR = "/var/www/filesystem"
var lan = make(map[int]interface{})

var users = make(map[int64]*User, 0)
var buckets = make(map[string]*Bucket, 0)

var mariadb *gorm.DB

func Init() {
	lan = l18n.ZH_LAN
	initDir()
	initTrans("zh")
	initDB()
	initData()
}

func initDir() {
	if os.Getenv("BaseDir") != "" {
		BASE_DIR = os.Getenv("BaseDir")
	}
	os.MkdirAll(BASE_DIR, os.ModePerm)
}

func initTrans(locale string) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLan := en.New()
		zhLan := zh.New()
		uni := ut.New(enLan, zhLan, enLan)
		trans, _ = uni.GetTranslator(locale) // locale 通常取决于 http 请求头的 'Accept-Language'
		switch locale {                      // 注册翻译器
		case "zh":
			zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			enTranslations.RegisterDefaultTranslations(v, trans)
		}
	}
}

func initDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/filesystem?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mariadb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("mysql connect error: " + err.Error())
	}
	mariadb.AutoMigrate(&User{}, &Bucket{})
}

func writeMem[T User | Bucket]() (user *User) {
	res := make([]T, 0)
	mariadb.Find(&res)
	for _, valRow := range res {
		switch val := any(valRow).(type) {
		case User:
			users[val.Id] = &val
			if val.UType == UTypeSystem && user == nil {
				user = &val
			}
		case Bucket:
			buckets[val.Name] = &val
		}
	}
	return
}

func initData() {
	if !checkExistUser("system") {
		auth := transform(randStringRunes(10))
		mariadb.Create(&User{
			Name:  "system",
			Auth:  auth,
			UType: UTypeSystem,
		})
	}
	systemuser := writeMem[User]()
	writeMem[Bucket]()
	if systemuser != nil {
		fmt.Printf("\tsystem user\n")
		fmt.Printf("\tuser: %s\n", systemuser.Name)
		fmt.Printf("\tauth: %s\n", systemuser.Auth)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func transform(s string) string {
	w := md5.New()
	w.Write([]byte(s))
	return hex.EncodeToString(w.Sum(nil))
}
