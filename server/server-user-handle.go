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
	"net/http"

	"github.com/book-library-seat-system/go-server/entity/seat"

	"github.com/book-library-seat-system/go-server/entity/user"
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

// testGET 测试GET
func testGET(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)

		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{})
	}
}

// testPOST 测试POST
func testPOST(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("Inter Post!")
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)

		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{})
	}
}

// createStudentHandle 创建一个新的用户
func createStudentHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("createStudentHandle!")
		// 解析参数
		param := parseReq(r)
		user.RegisterStudent(param["openID"], param["netID"], param["password"], param["school"])
		// 发回json
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// listStudentInfoHandle 显示用户信息
func listStudentInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("listStudentInfoHandle!")
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		pitem := user.GetStudent(param["openID"])
		// 解析json
		formatter.JSON(w, http.StatusOK, StudentRtnJson{
			School:    pitem.School,
			Violation: pitem.Violation,
			Seatinfos: seat.GetSeatinfoByStudentID(pitem.School, pitem.ID),
		})
	}
}
