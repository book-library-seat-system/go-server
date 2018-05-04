/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 常用的工具函数包
Date: 2018年5月4日 星期五 下午1:08
****************************************************************************/

package util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var logDivPath = "src/github.com/book-library-seat-system/go-server/log"
var logFilePath = "/" + time.Now().Format("2006-01-02") + ".txt"

func init() {
	//logDivPath = filepath.Join(*GetGOPATH(), logDivPath)
}

//GetGOPATH 获得用户环境的gopath
func GetGOPATH() *string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	goPath := strings.Split(os.Getenv("GOPATH"), sp)
	for _, v := range goPath {
		if _, err := os.Stat(filepath.Join(v, "/src/github.com/book-library-seat-system/go-server/util/util.go")); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}

func getFileHandle() *os.File {
	if _, err := os.Open(logDivPath + logFilePath); err != nil {
		os.Create(logDivPath + logFilePath)
	}

	// 以追加模式打开文件,并向文件写入
	fi, _ := os.OpenFile(logDivPath+logFilePath, os.O_RDWR|os.O_APPEND, 0)
	return fi
}

// CheckErr panic错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
