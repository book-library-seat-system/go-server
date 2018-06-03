/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 每一个user-handle层函数的测试处理
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server 

import (
	
	"net/url"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/book-library-seat-system/go-server/util"
)

func init() {
	ser := NewServer()
	ser.Run(":8899")
}

//测试testGet
func TesttestGET(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 发送http get请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8899/v1/test", strings.NewReader(""))
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

//测试testPost
func TesttestPost(t *testing.T){
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 发送http get请求
	client := &http.Client{}
	postValues := url.Values{}
	postValues.Add("school","sysu")
	postValues.Add("netID","15331116")
	resp, err := client.PostForm("http://localhost:8899/v1/test",postValues)
	CheckErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    CheckErr(err)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}
}

//测试createStudentHandle
func TestcreateStudentHandle(t *testing.T){

	http.HandleFunc("/createStudentCheck", createStudentHandle)
	//创建请求
	req, err := http.NewRequest("GET","/createStudentCheck", nil)
	if err != nil {
        t.Fatal(err)
    }
	//记录响应
	rr := httptest.NewRecorder()

	//检测返回状态码
	createStudentHandle(rr, req)
    if status := rr.Code; status != http.StatusOK{
		panic(errors.New("创建用户不正确"))
	}
	
}

//测试listStudentInfoHandle
func TestlistStudentInfoHandle(t *testing.T){
    defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	
	
	// 发送http get请求
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