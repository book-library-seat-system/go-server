/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 总服务端处理，路由设置
Date: 2018年5月4日 星期五 下午1:12
****************************************************************************/

package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer 新建客户端
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

// 初始化路由，分别初始化User部分和Meeting部分
func initRoutes(mx *mux.Router, formatter *render.Render) {
	initTestRoutes(mx, formatter)
	initUserRoutes(mx, formatter)
	initSeatRoute(mx, formatter)
}

// 用于测试部分
func initTestRoutes(mx *mux.Router, formatter *render.Render) {
	// 测试get
	mx.HandleFunc("/v1/test", testGET(formatter)).Methods("GET")
	// 测试post
	mx.HandleFunc("/v1/test", testPOST(formatter)).Methods("GET")
}

// 用户部分
func initUserRoutes(mx *mux.Router, formatter *render.Render) {
	// 创建用户
	mx.HandleFunc("/v1/users", createStudentHandle(formatter)).Methods("POST")
	// 显示用户信息
	mx.HandleFunc("/v1/users", listStudentInfoHandle(formatter)).Methods("GET")
}

//会议逻辑，路由设置
func initSeatRoute(mx *mux.Router, formatter *render.Render) {
	// 查看时间段
	mx.HandleFunc("/v1/timeintervals", showTimeIntervalInfoHandle(formatter)).Methods("GET")
	// 查看座位信息
	mx.HandleFunc("/v1/seats", showSeatInfoHandle(formatter)).Methods("GET")
	// 预约座位
	mx.HandleFunc("/v1/seat/book", bookSeatHandle(formatter)).Methods("POST")
	// 取消预约座位
	mx.HandleFunc("/v1/seat/unbook", unbookSeatHandle(formatter)).Methods("POST")
	// 签到座位
	mx.HandleFunc("/v1/seat/signin", signinSeatHandle(formatter)).Methods("POST")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
