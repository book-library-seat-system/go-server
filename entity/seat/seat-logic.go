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

// isInSigninTime 判断是否在签到时间，如果不是，抛出错误
func isInSigninTime(t time.Time) TimeInterval {
	// 如果不在50 - 60, 0 - 30时间段内，则不允许签到
	signintime := t.Add(10 * time.Minute)
	timeinterval := getCurrentTimeInterval(signintime)
	if signintime.Minute() < 40 || !timeinterval.Valid() {
		CheckErr(errors.New("108|签到时间不符"))
	}
	return timeinterval
}

// isInBookTime 判断是否在预定/取消预定时间，如果不是，抛出错误
func isInBookTime(t time.Time, timeinterval *TimeInterval) {
	// 如果不提前15min，则不允许预约
	booktime := t.Add(15 * time.Minute)
	if booktime.After(timeinterval.Begintime) {
		CheckErr(errors.New("109|预约/取消预约时间不符"))
	}
}

// insertOne 将一个seatinfo插入seatinfo数组，如果数组中有前一个时间段的，则接在后面
func insertOne(seatinfos []SeatInfo, newitem *SeatInfo) {
	for i := 0; i < len(seatinfos); i++ {
		if seatinfos[i].SeatID == newitem.SeatID &&
			seatinfos[i].Endtime == newitem.Begintime {
			seatinfos[i].Endtime = newitem.Endtime
			return
		}
	}
	seatinfos = append(seatinfos, *newitem)
}

// updateAllAfterSeat 从某一时间段起一直更新f函数为真的信息
func updateAllAfterSeat(school string, timeinterval TimeInterval, seatid int, f func(*Item) bool) {
	item := service.FindOneSeat(school, timeinterval, seatid)
	for ; f(&item); item = service.FindOneSeat(school, timeinterval, seatid) {
		service.UpdateOneSeat(school, timeinterval, item)
		timeinterval.AddOneHour()
	}
}

// unbookAllAfterSeat 对接下来的所有时间段，Book -> UnBook
func unbookAllAfterSeat(school string, timeinterval TimeInterval, seatid int, studentid string) {
	updateAllAfterSeat(school, timeinterval, seatid, func(item *Item) bool {
		if item.StudentID == studentid && item.Seatinfo == Book {
			item.StudentID = ""
			item.Seatinfo = UnBook
			return true
		}
		return false
	})
}

// signinAllAfterSeat 对接下来的所有时间段，Book -> Signin
func signinAllAfterSeat(school string, timeinterval TimeInterval, seatid int, studentid string) {
	updateAllAfterSeat(school, timeinterval, seatid, func(item *Item) bool {
		if item.StudentID == studentid && item.Seatinfo == Book {
			item.Seatinfo = Signin
			return true
		}
		return false
	})
}

// signoutAllAfterSeat 对接下来的所有时间段，Signin -> Signout
func signoutAllAfterSeat(school string, timeinterval TimeInterval, seatid int, studentid string) {
	updateAllAfterSeat(school, timeinterval, seatid, func(item *Item) bool {
		if item.StudentID == studentid && item.Seatinfo == Signin {
			item.Seatinfo = Signout
			return true
		}
		return false
	})
}

/*************************************************
Function: GetAllTimeInterval
Description: 得到允许预定的时间间隔（默认为两天）
InputParameter:
	school: 所查询的学校名字
Return: 可用时间间隔数组，以一小时为单位
*************************************************/
func GetAllTimeInterval(school string) []TimeInterval {
	// 预定开始于30min后的座位
	// 预定只允许今明两天的座位
	nowtimeinterval := getCurrentTimeInterval(time.Now().Add(30 * time.Minute))
	endday := nowtimeinterval.Begintime.Add(2 * 24 * time.Hour)
	nowtimeinterval.Endtime = time.Date(endday.Year(), endday.Month(), endday.Day(), 0, 0, 0, 0, endday.Location())
	return splitTimeInterval(nowtimeinterval)
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
		if items[i].Seatinfo == UnBook {
			count++
		}
	}
	return count
}

/*************************************************
Function: GetSeatinfoByStudentID
Description: 得到某个用户的所有预定的座位信息
InputParameter:
	school: 所查询的学校名字
	studentid: 学生id
Return: SeatInfo数组（时间是连续的）
*************************************************/
func GetBookSeatinfoByStudentID(school string, studentid string) []SeatInfo {
	seatinfos := service.FindBySchoolAndStudentID(school, studentid, Book)
	mergeseatinfos := []SeatInfo{}
	for _, seatinfo := range seatinfos {
		insertOne(mergeseatinfos, &seatinfo)
	}
	return mergeseatinfos
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
	isInBookTime(time.Now(), &timeinterval)

	validtimeintervals := splitTimeInterval(timeinterval)
	items := make([]Item, len(validtimeintervals))
	for i, validtimeinterval := range validtimeintervals {
		items[i] = service.FindOneSeat(school, validtimeinterval, seatid)
		if items[i].Seatinfo != UnBook {
			CheckErr(errors.New("106|该座位状态不符合要求"))
		}
	}
	for _, item := range items {
		item.StudentID = studentid
		item.Seatinfo = Book
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
	isInBookTime(time.Now(), &timeinterval)

	validtimeintervals := splitTimeInterval(timeinterval)
	items := make([]Item, len(validtimeintervals))
	for i, validtimeinterval := range validtimeintervals {
		items[i] = service.FindOneSeat(school, validtimeinterval, seatid)
		if items[i].Seatinfo != Book {
			CheckErr(errors.New("106|该座位状态不符合要求"))
		}
		if items[i].StudentID != studentid {
			CheckErr(errors.New("107|学生信息与该座位不符"))
		}
	}
	for _, item := range items {
		item.StudentID = ""
		item.Seatinfo = UnBook
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
	timeinterval := isInSigninTime(time.Now())

	// 测试座位状态
	item := service.FindOneSeat(school, timeinterval, seatid)
	if item.Seatinfo != Book {
		CheckErr(errors.New("106|该座位状态不符合要求"))
	}
	if item.StudentID != studentid {
		CheckErr(errors.New("107|学生信息与该座位不符"))
	}

	// 进行签到
	item.Seatinfo = Signin
	service.UpdateOneSeat(school, timeinterval, item)
	signinAllAfterSeat(school, *timeinterval.AddOneHour(), seatid, studentid)
}
