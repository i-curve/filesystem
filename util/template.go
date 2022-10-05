package util

import (
	"os"
	"path"
)

// 三元表达式
func TernaryExpr[T any](flag bool, a, b T) T {
	if flag {
		return a
	} else {
		return b
	}
}

// 创建嵌套目录
func MakDir(filename string) error {
	dirname := path.Dir("/data" + filename)
	return os.MkdirAll(dirname, 0666)
}

// 递归删除空目录
func DelDir(filename string) error {
	dirname := path.Dir(filename)
	if len(dirname) <= 5 {
		return nil
	}

	dirs, err := os.ReadDir(dirname)
	if err != nil || len(dirs) != 0 {
		return err
	} else if err = os.Remove(dirname); err != nil {
		return err
	} else {
		return DelDir(dirname)
	}
}
