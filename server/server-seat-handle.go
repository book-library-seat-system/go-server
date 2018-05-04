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

	"github.com/unrolled/render"
)

type resjson struct {
	Information string
}

//meetingjson 创建会议 存放json解析后的数据
type meetingjson struct {
	//会议主题
	Title string
	//会议参与者
	Participator []string
	//开始时间
	StartTime string
	//结束时间
	EndTime string
}

//增加会议参与者 存放json解析后的数据
type meetingAddjson struct {
	//会议参与者
	Participator []string
}

// 将参与者名字的类型[]string转成string方便数据库存储
func getParticipatorsName(p []string) string {
	s := ";"
	for i := 0; i < len(p); i++ {
		s = s + p[i] + ";"
	}
	fmt.Println(s)
	return s
}

// 返回cookie中携带的Name字段
func getCurrentUserNameMeeting(r *http.Request) string {
	cookie, _ := r.Cookie("username")
	if cookie != nil {
		return cookie.Value
	} else {
		fmt.Println("cookie nil")
	}
	return "unknown"
}

//getResponseJson 构造http response的json
func getResponseJson(info string) resjson {
	return resjson{
		Information: info}
}

func showTimeIntervalInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func showSeatInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func bookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func unbookSeatHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// //创建会议 /v1/meetings
// func createMeetingHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("createMeetingHandler")
// 		body, _ := ioutil.ReadAll(r.Body)
// 		var meetingj meetingjson
// 		if err := json.Unmarshal(body, &meetingj); err == nil {
// 			starttime, _ := time.Parse("2006-01-02 15:04:05", meetingj.StartTime)
// 			endtime, _ := time.Parse("2006-01-02 15:04:05", meetingj.EndTime)
// 			fmt.Println(starttime, endtime)
// 			meeting := meeting.Meeting{Title: meetingj.Title, Host: getCurrentUserNameMeeting(r),
// 				Participator: getParticipatorsName(meetingj.Participator), StartTime: starttime, EndTime: endtime}
// 			fmt.Println(meeting)
// 			err := meetingService.CreateMeeting(meeting)
// 			var info string
// 			if err != nil {
// 				info = err.Error()
// 				formatter.JSON(w, http.StatusBadRequest, getResponseJson(info))
// 			} else {
// 				info = "create meeting succeed"
// 				formatter.JSON(w, http.StatusOK, getResponseJson(info))
// 			}
// 			fmt.Println(info)
// 		} else {
// 			fmt.Println(err)
// 		}
// 		return
// 	}
// }

// //增加会议参与者 /v1/meeting/{title}/adding-participators
// func addParticipatorsHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.ParseForm()
// 		url := mux.Vars(r)
// 		title := url["title"]
// 		body, _ := ioutil.ReadAll(r.Body)
// 		var meetinga meetingAddjson
// 		if err := json.Unmarshal(body, &meetinga); err == nil {
// 			err := meetingService.AddMeetingParticipators(title, meetinga.Participator)
// 			var info string
// 			if err != nil {
// 				info = err.Error()
// 				formatter.JSON(w, http.StatusBadRequest, getResponseJson(info))
// 			} else {
// 				info = "add participators succeed"
// 				formatter.JSON(w, http.StatusOK, getResponseJson(info))
// 			}
// 		} else {
// 			formatter.JSON(w, http.StatusBadRequest, getResponseJson(err.Error()))
// 		}
// 	}
// }

// //删除会议参与者 /v1/meeting/{title}/deleting-participators
// func deleteParticipatorsHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 	}
// }

// //查询会议 /v1/users/query-meeting{?starttime,endtime}
// func queryMeetingsHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("queryMeetingsHandler")
// 		r.ParseForm()
// 		stime := r.Form["starttime"][0]
// 		etime := r.Form["endtime"][0]
// 		starttime, _ := time.Parse("2006-01-02 15:04:05", stime)
// 		endtime, _ := time.Parse("2006-01-02 15:04:05", etime)
// 		fmt.Println(stime, etime)
// 		queryMeetingResult, err := meetingService.QueryMeetings(getCurrentUserNameMeeting(r), starttime, endtime)
// 		var info string
// 		if err != nil {
// 			info = err.Error()
// 			formatter.JSON(w, http.StatusBadRequest, getResponseJson(info))
// 		} else {
// 			info = queryMeetingResult
// 			formatter.JSON(w, http.StatusOK, getResponseJson(info))
// 		}
// 	}
// }

// //取消会议 /v1/users/cancel-a-meeting/{title}
// func cancelMeetingHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("cancelMeetingHandler")
// 		r.ParseForm()
// 		url := mux.Vars(r)
// 		title := url["title"]
// 		fmt.Println(title)
// 		err := meetingService.CancelMeeting(title)
// 		var info string
// 		if err != nil {
// 			info = err.Error()
// 			formatter.JSON(w, http.StatusBadRequest, getResponseJson(info))
// 		} else {
// 			info = "cancel meeting succeed"
// 			formatter.JSON(w, http.StatusOK, getResponseJson(info))
// 		}
// 		fmt.Println(info)
// 	}
// }

// //退出会议 /v1/users/quit-meeting/{title}
// func quitMeetingHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 	}
// }

// //清空会议 /v1/users/cancel-all-meeting
// func clearAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("clear all")
// 	}
// }
