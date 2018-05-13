/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user方面服务端处理，包含user方面路由函数
	PS：本层不涉及逻辑判断，逻辑判断在user包部分
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/book-library-seat-system/go-server/entity/user"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

// 用于返回的模板Json
type resj struct {
	// 包含userItem属性
	user.Item
	// 返回user查询列表
	Users []user.Item `json:",omitempty"`
	// 表示结果
	Information string
}

// error.toString
func toString(err interface{}) string {
	if err == nil {
		return ""
	}
	return fmt.Sprint(err)
}

// 标准response JSON，只包含Success和Result
func stdResj(err interface{}) resj {
	return resj{
		Information: toString(err)}
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

func praseCookie(r *http.Request) string {
	// 解析cookie
	cookie, _ := r.Cookie("username")
	if cookie != nil {
		return cookie.Value
	}
	return ""
}

// 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		formatter.JSON(w, 500, stdResj(err))
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
			Errorinformation: ""
		})
	}
}

// 登录用户
func loginStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		// 使用user函数
		js := praseJSON(r)
		loginname := praseCookie(r)
		pitem := user.LoginUser(
		js.Get("ID").MustString(),
		js.Get("password").MustString())

		// 如果成功登录，设置cookie???
		cookie := http.Cookie{
		name:   "username",
		Value:  pitem.ID,
		}
		http.SetCookie(w, &cookie)

		resjson := stdResj(nil)
		resjson.Item = *pitem
		formatter.JSON(w, http.StatusOK, Returnjson{
			Errorcode:        0,
			Errorinformation: "",
			resjson
		})
	}
}

// 登出用户
func logoutStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// defer errResponse(w, formatter)

		// loginname := praseCookie(r)
		// user.LogoutUser(loginname)

		formatter.JSON(w, http.StatusOK, stdResj(nil))
	}
}

// 显示所有用户
//func listStudentInfoHandle(formatter *render.Render) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
		// defer errResponse(w, formatter)

		// loginname := praseCookie(r)
		// items := user.ListUsers(loginname)

		// resjson := stdResj(nil)
		// resjson.Users = items
		// formatter.JSON(w, http.StatusOK, resjson)
	//}
//}

// // 删除已登录用户
// func deleteUserHandle(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		defer errResponse(w, formatter)

// 		loginname := praseCookie(r)
// 		user.DeleteUser(loginname)

// 		// 如果成功删除，设置cookie
// 		cookie := http.Cookie{
// 			Name:   "username",
// 			Path:   "/",
// 			MaxAge: -1}
// 		http.SetCookie(w, &cookie)
// 		formatter.JSON(w, http.StatusOK, stdResj(nil))
// 	}
// }

// func undefinedHandler(formatter *render.Render) http.HandlerFunc {
//
// 	return func(w http.ResponseWriter, req *http.Request) {
// 	}
// }
