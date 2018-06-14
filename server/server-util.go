/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 服务端结构，用于json返回
Date: 2018年5月14日 星期一 上午10:25
****************************************************************************/

package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/book-library-seat-system/go-server/entity/user"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

// ErrorRtnJson 包含错误码和错误信息
type ErrorRtnJson struct {
	// 错误码
	Errorcode int `json:"errorcode"`
	// 错误信息
	Errorinformation string `json:"errorinformation"`
}

// errResponse 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		fmt.Println(err)
		var rtn ErrorRtnJson
		rtn.Errorcode, rtn.Errorinformation = HandleError(err)
		formatter.JSON(w, 500, rtn)
	}
}

// CheckUserLogin 判断用户是否登陆，如果没有，抛出错误
func CheckUserLogin(param map[string]string) {
	if param["openID"] == "" || len(param["openID"]) != 28 {
		CheckErr(errors.New("7|用户当前未登陆"))
	}
}

// getBeginandEndTime 解析参数中的begintime和endtime，如果不存在，报错
func getBegintimeAndEndtime(param map[string]string) (time.Time, time.Time) {
	begintime, err := time.ParseInLocation("2006-01-02 15:04:05", param["begintime"], time.Now().Location())
	CheckNewErr(err, "204|解析url参数错误")
	endtime, err := time.ParseInLocation("2006-01-02 15:04:05", param["endtime"], time.Now().Location())
	CheckNewErr(err, "204|解析url参数错误")
	return begintime, endtime
}

// parseReq 解析参数，如果解析到openID参数，则另外加入school信息
func parseReq(r *http.Request) map[string]string {
	param := (map[string]string)(nil)
	if r.Method == "GET" {
		param = parseReqByGet(r)
	} else if r.Method == "POST" {
		param = parseReqByPost(r)
	}
	fmt.Println(param)
	if param != nil && param["openID"] != "" && param["school"] == "" {
		param["school"] = user.GetStudentsSchool(param["openID"])
	}
	return param
}

// parseReqByPost Post请求方式解析参数
func parseReqByPost(r *http.Request) map[string]string {
	// 解析参数
	err := r.ParseForm()
	CheckNewErr(err, "203|解析json错误")
	body, err := ioutil.ReadAll(r.Body)
	CheckNewErr(err, "203|解析json错误")
	defer r.Body.Close()

	// 解析json，转换成map
	temp, err := simplejson.NewJson(body)
	CheckNewErr(err, "203|解析json错误")
	m, _ := temp.Map()
	CheckNewErr(err, "203|解析json错误")
	rtnmap := make(map[string]string)
	for k, v := range m {
		rtnmap[k] = fmt.Sprint(v)
	}
	return rtnmap
}

// parseReqByGet Get请求方式解析参数
func parseReqByGet(r *http.Request) map[string]string {
	// 解析参数
	err := r.ParseForm()
	CheckNewErr(err, "204|解析url参数错误")

	// 解析ID
	rtnmap := make(map[string]string)
	for k, v := range r.Form {
		rtnmap[k] = v[0]
	}
	return rtnmap
}
