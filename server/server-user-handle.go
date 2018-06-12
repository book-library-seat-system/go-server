/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user方面服务端处理，包含user方面路由函数
	PS：本层不涉及逻辑判断，逻辑判断在user包部分
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/book-library-seat-system/go-server/entity/seat"

	"github.com/book-library-seat-system/go-server/entity/user"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

// UserReturnjson 用于返回student的模板Json
type StudentRtnJson struct {
	// 学生所在学校
	School string `json:"school,omitempty"`
	// 学生被警告次数
	Violation int `json:"violation,omitempty"`
	// 预约的座位信息数组
	Seatinfos []seat.SeatInfo `json:"bookseatinfos,omitempty"`
	// 包含错误信息
	ErrorRtnJson
}

// 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		fmt.Println(err)
		var rtn ErrorRtnJson
		rtn.Errorcode, rtn.Errorinformation = HandleError(err)
		formatter.JSON(w, 500, rtn)
	}
}

// testGET
func testGET(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		// 解析url参数
		param := parseUrl(r)
		if _, ok := param["openID"]; !ok {
			CheckErr(errors.New("7|用户当前未登陆"))
		}

		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{})
	}
}

func testPOST(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)

		fmt.Println("Inter Post!")
		// 解析url参数
		js := parseJSON(r)
		fmt.Println(*js)
		_, err := js.Get("openID").String()
		if err != nil {
			CheckErr(errors.New("7|用户当前未登陆"))
		}

		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{})
	}
}

// 创建一个新的用户
func createStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("createStudentHandle!")

		// 解析json数据
		js := parseJSON(r)
		fmt.Println(js)
		user.RegisterStudent(
			js.Get("openID").MustString(),
			js.Get("netID").MustString(),
			js.Get("password").MustString(),
			js.Get("school").MustString())

		// 发回json
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// 显示用户信息
func listStudentInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("listStudentInfoHandle!")

		// 解析url参数
		param := parseUrl(r)
		if _, ok := param["openID"]; !ok {
			CheckErr(errors.New("7|用户当前未登陆"))
		}

		pitem := user.GetStudent(param["openID"])
		fmt.Println(*pitem)
		seatinfos := seat.GetSeatinfoByStudentID(pitem.School, pitem.ID)
		fmt.Println(seatinfos)

		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{
			School:    pitem.School,
			Violation: pitem.Violation,
			Seatinfos: seatinfos,
		})
	}
}
