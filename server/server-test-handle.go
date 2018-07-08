/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 用于测试web请求的文件
Date: 2018年7月8日 星期日 下午2:02
****************************************************************************/

package server

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

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
