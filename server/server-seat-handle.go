/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat方面服务端处理，包含seat方面路由函数
	PS：本层不涉及逻辑判断，逻辑判断在seat包部分
Date: 2018年5月4日 星期五 下午1:13
****************************************************************************/

package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/book-library-seat-system/go-server/entity/seat"
	. "github.com/book-library-seat-system/go-server/util"
	"github.com/unrolled/render"
)

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

type SeatinfoRtnJson struct {
	// 座位信息
	Seatinfo []int `json:"seatinfo,omitempty"`
	// 错误信息
	ErrorRtnJson
}

func showTimeIntervalInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showTimeIntervalInfoHandle")

		// 解析cookie数据
		_, school, err := parseCookie(r)
		CheckNewErr(err, "7|用户当前未登陆")

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

func showSeatInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showSeatInfoHandle")

		// 解析json数据
		err := r.ParseForm()
		CheckNewErr(err, "203|解析json错误")
		begintime, err := time.Parse("2006-01-02 15:04:05", r.Form["begintime"][0])
		CheckNewErr(err, "203|解析json错误")
		endtime, err := time.Parse("2006-01-02 15:04:05", r.Form["endtime"][0])
		CheckNewErr(err, "203|解析json错误")

		// 解析cookie数据
		_, school, err := parseCookie(r)
		CheckNewErr(err, "7|用户当前未登陆")

		// 从数据库得到数据
		rtnjson := SeatinfoRtnJson{
			Seatinfo: seat.GetAllSeatinfo(school, seat.TimeInterval{begintime, endtime}),
		}

		// 发送json
		formatter.JSON(w, http.StatusOK, rtnjson)
	}
}

func bookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("bookSeatHandle")

		// 解析json
		js := parseJSON(r)
		begintime, err := time.Parse("2006-01-02 15:04:05", js.Get("begintime").MustString())
		CheckNewErr(err, "203|解析json错误")
		endtime, err := time.Parse("2006-01-02 15:04:05", js.Get("endtime").MustString())
		CheckNewErr(err, "203|解析json错误")

		// 解析cookie数据
		studentid, school, err := parseCookie(r)
		CheckNewErr(err, "7|用户当前未登陆")

		// 进行预约
		seat.BookSeat(school, seat.TimeInterval{begintime, endtime}, studentid, js.Get("seatID").MustInt())
		formatter.JSON(w, http.StatusOK, nil)
	}
}

func unbookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("unbookSeatHandle")

		// 解析json
		js := parseJSON(r)
		begintime, err := time.Parse("2006-01-02 15:04:05", js.Get("begintime").MustString())
		CheckNewErr(err, "203|解析json错误")
		endtime, err := time.Parse("2006-01-02 15:04:05", js.Get("endtime").MustString())
		CheckNewErr(err, "203|解析json错误")

		// 解析cookie数据
		studentid, school, err := parseCookie(r)
		CheckNewErr(err, "7|用户当前未登陆")

		// 进行预约
		seat.UnbookSeat(school, seat.TimeInterval{begintime, endtime}, studentid, js.Get("seatID").MustInt())
		formatter.JSON(w, http.StatusOK, nil)
	}
}
