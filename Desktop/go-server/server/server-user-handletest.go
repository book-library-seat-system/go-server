/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 每一个user-handle层函数的测试处理
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
func TesttestPost(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	data := make(url.Values)
	data["school"] = []string{"sysu"}
	data["netID"] = []string{"15331116"}
	res, err := http.PostForm("http://localhost:8899/v1/test", url.Values{"school":{"sysu"},"netId":{"15331116"}})
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(body)
	CheckErr(err)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}
}

//测试createStudentHandle
func TestcreateStudentHandle(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	data := make(url.Values)
	data["school"] = []string{"sysu"}
	data["netID"] = []string{"15331116"}
	res, err := http.PostForm("http://localhost:8899/v1/test", url.Values{"school":{"sysu"},"netId":{"15331116"}})
	
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(body)
	CheckErr(err)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}

}

//测试listStudentInfoHandle
func TestlistStudentInfoHandle(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 发送http get请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8899/v1/users", nil)
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
	fmt.Println(body)

	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		panic(errors.New("返回错误不正确"))
	}

}
