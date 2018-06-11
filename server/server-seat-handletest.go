/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 每一个seat-handle层函数的测试处理
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/
package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/book-library-seat-system/go-server/util"
)

func testShowTimeIntervalInfoHandle(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	//发送http请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8899/v1/test", nil)
	CheckErr(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()

    // 接收响应并且读取body信息
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 103 || errresp.Errorinformation != "数据库座位信息查找出现错误" {
		panic(errors.New("返回错误不正确"))
	}
}

func testShowSeatInfoHandle(t *testing.T) {
    // defer func() {
	// 	if err := recover(); err != nil {
	// 		t.Error(err)
	// 	}
	// }()

	//发送http请求
}

func testBookSeatHandle(t *testing.T) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		t.Error(err)
	// 	}
	// }()

	//发送http请求
}

func testUnbookSeatHandle(t *testing.T) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		t.Error(err)
	// 	}
	// }()

	//发送http请求
}

func testSigninSeatHandle(t *testing.T) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		t.Error(err)
	// 	}
	// }()

	//发送http请求
}
