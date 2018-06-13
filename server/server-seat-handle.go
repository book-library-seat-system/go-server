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
	Seatinfos []int `json:"seatinfos,omitempty"`
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
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		// 从数据库获取数据
		timeintervals := seat.GetAllTimeInterval(param["school"])
		rtnjson := TimeintervalRtnJson{}
		for i := 0; i < len(timeintervals); i++ {
			rtnjson.Timeintervals = append(rtnjson.Timeintervals, TimeintervalJson{
				TimeInterval: timeintervals[i],
				Restseatsnum: len(seat.GetAllSeatinfo(param["school"], timeintervals[i])),
			})
		}
		formatter.JSON(w, http.StatusOK, rtnjson)
	}
}

// showSeatInfoHandle 返回座位信息
func showSeatInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showSeatInfoHandle")
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		begintime, endtime := getBegintimeAndEndtime(param)
		// 从数据库得到数据
		rtnjson := SeatinfoRtnJson{
			Seatinfos: seat.GetAllSeatinfo(param["school"], seat.TimeInterval{begintime, endtime}),
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
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		if user.GetStudent(param["openID"]).IsPunished() {
			CheckErr(errors.New("110|用户当前被惩罚"))
		}
		begintime, endtime := getBegintimeAndEndtime(param)
		// 进行预约
		seat.BookSeat(param["school"], seat.TimeInterval{begintime, endtime}, param["openID"], String2Int(param["seatID"]))
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// unbookSeatHandle 取消预约座位
func unbookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("unbookSeatHandle")
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		begintime, endtime := getBegintimeAndEndtime(param)
		// 进行预约
		seat.UnbookSeat(param["school"], seat.TimeInterval{begintime, endtime}, param["openID"], String2Int(param["seatID"]))
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}

// signinSeatHandle 签到座位
func signinSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("signinSeatHandle")
		// 解析参数
		param := parseReq(r)
		CheckUserLogin(param)
		// 进行签到
		seat.SigninSeat(param["school"], param["openID"], String2Int(param["seatID"]))
		formatter.JSON(w, http.StatusOK, ErrorRtnJson{})
	}
}
