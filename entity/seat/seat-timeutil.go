/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 关于时间的处理函数
Date: 2018年5月30日 星期三 上午9:55
****************************************************************************/

package seat

import (
	"strconv"
	"time"

	. "github.com/book-library-seat-system/go-server/util"
)

func getToday(t time.Time, hour int, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, 0, 0, t.Location())
}

/*************************************************
Function: getCurrentTimeInterval
Description: 得到给定时间所处的时间段
InputParameter:
	t: 给定时间
Return: 时间段
*************************************************/
func getCurrentTimeInterval(t time.Time) TimeInterval {
	begintime := getToday(t, t.Hour(), 0)
	endtime := begintime.Add(time.Hour)
	return TimeInterval{begintime, endtime}
}

// 将TimeInterval划分成1h间隔
func splitTimeInterval(timeinterval TimeInterval) []TimeInterval {
	rtntimeintervals := []TimeInterval{}
	for btime := timeinterval.Begintime; btime != timeinterval.Endtime; btime = btime.Add(time.Hour) {
		etime := btime.Add(time.Hour)
		newtimeinterval := TimeInterval{btime, etime}
		if newtimeinterval.Valid() {
			rtntimeintervals = append(rtntimeintervals, newtimeinterval)
		}
	}
	return rtntimeintervals
}

// 通过配置文件，读取有效时间段
func currentTimeIntervals() []TimeInterval {
	daysstr := ReadFromIniFile("TimeInterval", "days")
	days, _ := strconv.Atoi(daysstr)

	// 生成时间段
	nowtimeinterval := getCurrentTimeInterval(time.Now())
	endtime := nowtimeinterval.Begintime
	for i := 0; i < days; i++ {
		endtime = endtime.Add(24 * time.Hour)
	}
	nowtimeinterval.Endtime = time.Date(endtime.Year(), endtime.Month(), endtime.Day(), 0, 0, 0, 0, endtime.Location())
	return splitTimeInterval(nowtimeinterval)
}

// Valid 判断TimeInterval时间段是否有效（8：00-22:00）
func (t *TimeInterval) Valid() bool {
	return Valid(t.Begintime)
}

// Valid 判断Time时间段是否有效（8：00-22:00）
func Valid(t time.Time) bool {
	return t.Hour() >= 8 && t.Hour() < 22
}

// AddOneHour 时间相加1小时
func (t *TimeInterval) AddOneHour() *TimeInterval {
	t.Begintime = t.Begintime.Add(time.Hour)
	t.Endtime = t.Endtime.Add(time.Hour)
	return t
}

// Add 时间相加（本身不相加）
func (t *TimeInterval) Add(d time.Duration) TimeInterval {
	return TimeInterval{t.Begintime.Add(d), t.Endtime.Add(d)}
}

// Equal TimeInterval相等比较
func (t1 *TimeInterval) Equal(t2 TimeInterval) bool {
	return t1.Begintime.Equal(t2.Begintime) &&
		t2.Endtime.Equal(t2.Endtime)
}
