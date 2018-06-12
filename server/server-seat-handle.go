/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat方面服务端处理，包含seat方面路由函数
	PS：本层不涉及逻辑判断，逻辑判断在seat包部分
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/book-library-seat-system/go-server/entity/seat"
	"github.com/book-library-seat-system/go-server/entity/user"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

// TimeIntervalJson 在TimeInterval的前提下新加入一个字段
type TimeintervalJson struct {
	// 时间段
	seat.TimeInterval
	// 剩余座位数量
	Restseatsnum int `json:"restseatsnum"`
}

// TimeintervalRtnJson 返回时间戳数组
type TimeintervalRtnJson struct {
	// 时间段信息
	Timeintervals []TimeintervalJson `json:"timeintervals,omitempty"`
	// 错误信息
	ErrorRtnJson
}

// SeatinfoRtnJson 返回的座位信息数组
type SeatinfoRtnJson struct {
	// 座位信息
	Seatinfos []bool `json:"seatinfos,omitempty"`
	// 错误信息
	ErrorRtnJson
}

// BookedSeatinfoRtnJson 返回的座位信息数组
type BookedSeatinfoRtnJson struct {
	// 座位信息
	Seatinfos []seat.SeatInfo `json:"seatinfos,omitempty"`
	// 错误信息
	ErrorRtnJson
}

// showTimeIntervalInfoHandle 返回时间段信息
func showTimeIntervalInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showTimeIntervalInfoHandle")

		// 解析url参数
		param := parseUrl(r)
		if _, ok := param["openID"]; !ok {
			CheckErr(errors.New("7|用户当前未登陆"))
		}

		// 查询学校
		school := user.GetStudentsSchool(param["openID"])

		// 从数据库获取数据
		timeintervals := seat.GetAllTimeInterval(school)
		rtnjson := TimeintervalRtnJson{}
		for i := 0; i < len(timeintervals); i++ {
			rtnjson.Timeintervals = append(rtnjson.Timeintervals, TimeintervalJson{
				TimeInterval: timeintervals[i],
				Restseatsnum: len(seat.GetAllSeatinfo(school, timeintervals[i])),
			})
		}

		// 发送json
		formatter.JSON(w, http.StatusOK, rtnjson)
	}
}

// showSeatInfoHandle 返回座位信息
func showSeatInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showSeatInfoHandle")

		// 解析url参数数据
		param := parseUrl(r)
		if _, ok := param["openID"]; !ok {
			CheckErr(errors.New("7|用户当前未登陆"))
		}
		school := user.GetStudentsSchool(param["openID"])
		begintime, err := time.ParseInLocation("2006-01-02 15:04:05", param["begintime"], time.Now().Location())
		CheckNewErr(err, "204|解析url参数错误")
		endtime, err := time.ParseInLocation("2006-01-02 15:04:05", param["endtime"], time.Now().Location())
		CheckNewErr(err, "204|解析url参数错误")

		// 从数据库得到数据
		rtnjson := SeatinfoRtnJson{
			Seatinfos: seat.GetAllSeatinfo(school, seat.TimeInterval{begintime, endtime}),
		}

		// 发送json
		formatter.JSON(w, http.StatusOK, rtnjson)
	}
}

// bookSeatHandle 预约座位
func bookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("bookSeatHandle")

		// 解析json
		js := parseJSON(r)
		student := user.GetStudent(js.Get("openID").MustString())
		if student.IsPunished() {
			CheckErr(errors.New("110|用户当前被惩罚"))
		}
		begintime, err := time.ParseInLocation("2006-01-02 15:04:05", js.Get("begintime").MustString(), time.Now().Location())
		CheckNewErr(err, "203|解析json错误")
		endtime, err := time.ParseInLocation("2006-01-02 15:04:05", js.Get("endtime").MustString(), time.Now().Location())
		CheckNewErr(err, "203|解析json错误")

		// 进行预约
		seat.BookSeat(student.School, seat.TimeInterval{begintime, endtime}, student.ID, js.Get("seatID").MustInt())
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// unbookSeatHandle 取消预约座位
func unbookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("unbookSeatHandle")

		// 解析json
		js := parseJSON(r)
		studentid := js.Get("openID").MustString()
		school := user.GetStudentsSchool(studentid)
		begintime, err := time.ParseInLocation("2006-01-02 15:04:05", js.Get("begintime").MustString(), time.Now().Location())
		CheckNewErr(err, "203|解析json错误")
		endtime, err := time.ParseInLocation("2006-01-02 15:04:05", js.Get("endtime").MustString(), time.Now().Location())
		CheckNewErr(err, "203|解析json错误")

		// 进行预约
		seat.UnbookSeat(school, seat.TimeInterval{begintime, endtime}, studentid, js.Get("seatID").MustInt())
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// signinSeatHandle 签到座位
func signinSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("signinSeatHandle")

		// 解析json
		js := parseJSON(r)
		studentid := js.Get("openID").MustString()
		school := user.GetStudentsSchool(studentid)

		// 进行签到
		seat.SigninSeat(school, studentid, js.Get("seatID").MustInt())
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}
