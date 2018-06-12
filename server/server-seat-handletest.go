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
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/book-library-seat-system/go-server/util"
)

func TestShowTimeIntervalInfoHandle(t *testing.T) {
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
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}
}

func TestShowSeatInfoHandle1(t *testing.T) {
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
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}
}

func TestShowSeatInfoHandle2(t *testing.T) {
    defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// //发送http请求
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "http://localhost:8899/v1/test", nil)
	// CheckErr(err)
	// req.Header.Set("Content-Type", "application/json")
	// resp, err := client.Do(req)
	// CheckErr(err)
	// defer resp.Body.Close()

    // // 接收响应并且读取body信息
	// body, err := ioutil.ReadAll(resp.Body)
	// CheckErr(err)
	// errresp := ErrorRtnJson{}
	// json.Unmarshal(body, &errresp)

	// // 判断返回的错误信息是否符合要求
	// if errresp.Errorcode != 204 || errresp.Errorinformation != "解析url参数错误" {
	// 	panic(errors.New("返回错误不正确"))
	// }
}

func testBookSeatHandle(t *testing.T) {
	
}

func testUnbookSeatHandle(t *testing.T) {
	
}

func testSigninSeatHandle(t *testing.T) {
	
}
