/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的逻辑层，使用dao层接口，为上层控制层（server层）提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:52
****************************************************************************/

package seat

import (
	"errors"
	"time"

	. "github.com/book-library-seat-system/go-server/util"
)

var userItemsFilePath = "src/github.com/book-library-seat-system/go-server/orm/UserItems.json"
var currentUserFilePath = "src/github.com/book-library-seat-system/go-server/orm/Current.txt"

/*************************************************
Function: GetAllTimeInterval
Description: 得到所有的时间间隔
InputParameter:
	school: 所查询的学校名字
Return: 可用时间间隔数组，以一小时为单位
*************************************************/
func GetAllTimeInterval(school string) []TimeInterval {
	titems := service.FindBySchool(school)
	timeintervals := []TimeInterval{}
	for i := 0; i < len(titems); i++ {
		timeintervals[i] = titems[i].Timeinterval
	}
	return timeintervals
}

/*************************************************
Function: GetAllSeatinfo
Description: 得到某时间段所有座位的信息，数组下标代表位置
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
Return: 该时间段的座位预约信息，用int数组保存
*************************************************/
func GetAllSeatinfo(school string, timeinterval TimeInterval) []int {
	items := service.FindBySchoolAndTimeInterval(school, timeinterval)
	seatinfo := make([]int, len(items))
	for i, item := range items {
		seatinfo[i] = item.Seatinfo
	}
	return seatinfo
}

/*************************************************
Function: GetAllUnbookSeatNumber
Description: 得到某时间段的未预约座位数量
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
Return: 未预约的座位数量
*************************************************/
func GetAllUnbookSeatNumber(school string, timeinterval TimeInterval) int {
	count := 0
	items := service.FindBySchoolAndTimeInterval(school, timeinterval)
	for i := 0; i < len(items); i++ {
		if items[i].Seatinfo == 0 {
			count++
		}
	}
	return count
}

/*************************************************
Function: BookSeat
Description: 预约座位
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳，以小时为单位
	studentid: 预约学生ID
	seatid: 座位ID，即数组下标
Return: none
*************************************************/
func BookSeat(school string, timeinterval TimeInterval, studentid string, seatid int) {
	validtimeintervals := splitTimeInterval(timeinterval)
	items := make([]Item, len(validtimeintervals))
	for i, validtimeinterval := range validtimeintervals {
		items[i] = service.FindOneSeat(school, validtimeinterval, seatid)
		if items[i].Seatinfo != 0 {
			CheckErr(errors.New("106|该座位状态不符合要求"))
		}
	}
	for _, item := range items {
		item.StudentID = studentid
		item.Seatinfo = 1
		service.UpdateOneSeat(school, timeinterval, item)
	}
}

/*************************************************
Function: UnbookSeat
Description: 取消预约座位
InputParameter:
	school: 所查询的学校名字
	timeinterval: 查询的时间戳
	studentid: 预约学生ID
	seatid: 座位ID，即数组下标
Return: none
*************************************************/
func UnbookSeat(school string, timeinterval TimeInterval, studentid string, seatid int) {
	validtimeintervals := splitTimeInterval(timeinterval)
	items := make([]Item, len(validtimeintervals))
	for i, validtimeinterval := range validtimeintervals {
		items[i] = service.FindOneSeat(school, validtimeinterval, seatid)
		if items[i].Seatinfo != 1 {
			CheckErr(errors.New("106|该座位状态不符合要求"))
		}
		if items[i].StudentID != studentid {
			CheckErr(errors.New("107|学生信息与该座位不符"))
		}
	}
	for _, item := range items {
		item.StudentID = ""
		item.Seatinfo = 0
		service.UpdateOneSeat(school, timeinterval, item)
	}
}

/*************************************************
Function: SigninSeat
Description: 签到座位
InputParameter:
	school: 所查询的学校名字
	studentid: 预约学生ID
	seatid: 座位ID，即数组下标
Return: none
*************************************************/
func SigninSeat(school string, studentid string, seatid int) {
	// 得到应该签到的时间
	m10, _ := time.ParseDuration("10m")
	signtime := time.Now().Add(m10)
	timeinterval := getCurrentTimeInterval(signtime)
	if signtime.Minute() > 30 || !timeinterval.Valid() {
		CheckErr(errors.New("108|签到时间不符"))
	}

	item := service.FindOneSeat(school, timeinterval, seatid)
	if item.Seatinfo != 1 {
		CheckErr(errors.New("106|该座位状态不符合要求"))
	}
	if item.StudentID != studentid {
		CheckErr(errors.New("107|学生信息与该座位不符"))
	}
	item.Seatinfo = 2
	service.UpdateOneSeat(school, timeinterval, item)

	for {
		h, _ := time.ParseDuration("1h")
		timeinterval.Begintime.Add(h)
		timeinterval.Endtime.Add(h)
		if !timeinterval.Valid() {
			break
		}
		item = service.FindOneSeat(school, timeinterval, seatid)
		if item.Seatinfo == 1 && item.StudentID == studentid {
			item.Seatinfo = 2
			service.UpdateOneSeat(school, timeinterval, item)
		} else {
			break
		}
	}
}
