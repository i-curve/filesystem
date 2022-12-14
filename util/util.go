package util

import (
	"errors"
	"filesystem/config"
	"filesystem/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/i-curve/coding"
)

// 文件上传
func WriteFile(base string, file *multipart.FileHeader) (*model.Reply, coding.Code) {
	if file == nil {
		return nil, coding.New(coding.StatusOK, 400, "上传文件不能为空")
	}
	base = "/" + TernaryExpr(base != "", base, file.Filename)
	filename := base
	MakDir(filename)
	f, _ := file.Open()
	bys, _ := io.ReadAll(f)
	err := os.WriteFile("/data"+filename, bys, 0666)
	return &model.Reply{URL: config.BaseURL + filename, ShortURL: base}, coding.New(coding.StatusOK, 500, err)
}

// 复制文件
func CopyFile(oldPath, newPath string) coding.Code {
	oldFilename := "/data/" + oldPath
	newFilename := "/data/" + newPath
	bys, err := os.ReadFile(oldFilename)
	if err != nil {
		coding.New(http.StatusOK, 400, "原始文件不存在")
	}
	err = os.WriteFile(newFilename, bys, 0666)
	return coding.New(http.StatusOK, 500, TernaryExpr(err == nil, nil, errors.New("文件无法存储")))
}

// 移动文件
func MoveFile(oldPath, newPath string) coding.Code {
	err := CopyFile(oldPath, newPath)
	if err != nil {
		return err
	}
	return DeleteFile(oldPath)
}

// 文件删除
func DeleteFile(base string) coding.Code {
	filename := "/data/" + base
	if _, err := os.Stat(filename); err != nil {
		return coding.New(coding.StatusOK, 400, "文件不存在:"+err.Error())
	}
	if err := os.Remove(filename); err != nil {
		coding.New(http.StatusOK, 400, "删除失败:"+err.Error())
	}
	return coding.New(http.StatusOK, 400, DelDir(filename))
}
