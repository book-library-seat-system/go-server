/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 每一个test-handle层函数的测试处理
Date: 2018年7月8日 星期日 下午2:07
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
)

// CheckErr panic错误
func CheckTErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

//测试testGet
func TesttestGET(t *testing.T) {
	// 发送http get请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8899/v1/test?openID=123%20123", strings.NewReader(""))
	CheckTErr(err, t)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	CheckTErr(err, t)
	defer resp.Body.Close()
	// 接收响应并且读取body信息
	body, err := ioutil.ReadAll(resp.Body)
	CheckTErr(err, t)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)
	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		t.Error(errors.New("返回错误不正确"))
	}
}

//测试testPost
func TesttestPost(t *testing.T) {
	// 发送http get请求
	client := &http.Client{}
	postValues := url.Values{}
	postValues.Add("school", "sysu")
	postValues.Add("netID", "15331116")
	resp, err := client.PostForm("http://localhost:8899/v1/test", postValues)
	CheckTErr(err, t)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckTErr(err, t)
	errresp := ErrorRtnJson{}
	json.Unmarshal(body, &errresp)
	fmt.Println(body)
	// 判断返回的错误信息是否符合要求
	if errresp.Errorcode != 7 || errresp.Errorinformation != "用户当前未登陆" {
		t.Error(errors.New("返回错误不正确"))
	}
}
