package handle

import (
	"os"
	"path"

	"github.com/i-curve/copier"
)

func copy[T any](sour any) *T {
	dest := new(T)
	copier.Copy(&dest, sour)
	return dest
}

// 创建嵌套目录
func mkDir(filename string) error {
	dirname := path.Dir("/data" + filename)
	return os.MkdirAll(dirname, 0666)
}

// 递归删除空目录
func delDir(filename string) error {
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
		return delDir(dirname)
	}
}

// // 移动文件
// func moveFile(oldPath, newPath string) coding.Code {
// 	err := CopyFile(oldPath, newPath)
// 	if err != nil {
// 		return err
// 	}
// 	return DeleteFile(oldPath)
// }
