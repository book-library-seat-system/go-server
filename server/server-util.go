/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 服务端结构，用于json返回
Date: 2018年5月14日 星期一 上午10:25
****************************************************************************/

package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	. "github.com/book-library-seat-system/go-server/util"
)

// ErrorRtnJson 包含错误码和错误信息
type ErrorRtnJson struct {
	// 错误码
	Errorcode int `json:"errorcode"`
	// 错误信息
	Errorinformation string `json:"errorinformation"`
}

// 解析传过来的JSON
func parseJSON(r *http.Request) *simplejson.Json {
	// 解析参数
	err := r.ParseForm()
	CheckNewErr(err, "203|解析json错误")

	// 解析json
	body, err := ioutil.ReadAll(r.Body)
	CheckNewErr(err, "203|解析json错误")
	defer r.Body.Close()

	temp, err := simplejson.NewJson(body)
	CheckNewErr(err, "203|解析json错误")
	return temp
}

// 解析传过来的Url
func parseUrl(r *http.Request) map[string]string {
	// 解析参数
	err := r.ParseForm()
	CheckNewErr(err, "204|解析url参数错误")

	// 解析ID
	rtnmap := make(map[string]string)
	for k, v := range r.Form {
		fmt.Println(k, ":", v[0])
		rtnmap[k] = v[0]
	}
	return rtnmap
}

// 得到设置的cookie
func getCookie(key string, value string) *http.Cookie {
	return &http.Cookie{
		Name:   key,
		Value:  value,
		Path:   "/",
		Domain: "localhost",
		MaxAge: 3600,
	}
}
