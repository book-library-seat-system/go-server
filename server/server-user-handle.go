/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user方面服务端处理，包含user方面路由函数
	PS：本层不涉及逻辑判断，逻辑判断在user包部分
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/book-library-seat-system/go-server/entity/user"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

// 用于返回的模板Json
type Returnjson struct {
	// 包含userItem属性
	user.Item
	// 错误码
	Errorcode int `json:"errorcode,omitempty"`
	// 错误信息
	Errorinformation string `json:"errorinformation,omitempty"`
}

// 解析传过来的JSON和cookie
func praseJSON(r *http.Request) *simplejson.Json {
	// 解析json
	body, err := ioutil.ReadAll(r.Body)
	CheckErr(err)
	defer r.Body.Close()

	temp, err := simplejson.NewJson(body)
	CheckErr(err)
	return temp
}

// func praseCookie(r *http.Request) string {
// 	// 解析cookie
// 	cookie, _ := r.Cookie("username")
// 	if cookie != nil {
// 		return cookie.Value
// 	}
// 	return ""
// }

// 返回json信息
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		var rtn Returnjson
		errstrs := strings.Split(err.(error).Error(), "|")
		if len(errstrs) != 2 {
			rtn.Errorcode = 200
			rtn.Errorinformation = err.(error).Error()
		} else if rtn.Errorcode, err = strconv.Atoi(errstrs[0]); err != nil {
			rtn.Errorcode = 200
			rtn.Errorinformation = "未定义错误"
		}
		rtn.Errorinformation = errstrs[1]

		// 发送json返回
		formatter.JSON(w, 500, rtn)
	}
}

// test
func test(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// 创建一个新的用户
func createStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		// 解析json数据
		js := praseJSON(r)
		user.RegisterStudent(
			js.Get("ID").MustString(),
			js.Get("name").MustString(),
			js.Get("password").MustString(),
			js.Get("email").MustString(),
			js.Get("school").MustString())

		formatter.JSON(w, http.StatusOK, Returnjson{
			Errorcode:        0,
			Errorinformation: "",
		})
	}
}

// 登录用户
func loginStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// defer errResponse(w, formatter)

		// // 使用user函数
		// js := praseJSON(r)
		// loginname := praseCookie(r)
		// pitem := user.LoginUser(
		// 	js.Get("Name").MustString(),
		// 	js.Get("Password").MustString(),
		// 	loginname)

		// // 如果成功登录，设置cookie
		// cookie := http.Cookie{
		// 	Name:   "username",
		// 	Value:  pitem.Name,
		// 	Path:   "/",
		// 	MaxAge: 1200}
		// http.SetCookie(w, &cookie)

		// resjson := stdResj(nil)
		// resjson.Item = *pitem
		// formatter.JSON(w, http.StatusOK, resjson)
	}
}

// 登出用户
func logoutStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// defer errResponse(w, formatter)

		// loginname := praseCookie(r)
		// user.LogoutUser(loginname)

		// formatter.JSON(w, http.StatusOK, stdResj(nil))
	}
}

// 显示所有用户
func listStudentInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// defer errResponse(w, formatter)

		// loginname := praseCookie(r)
		// items := user.ListUsers(loginname)

		// resjson := stdResj(nil)
		// resjson.Users = items
		// formatter.JSON(w, http.StatusOK, resjson)
	}
}
