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

/*************************************************
Function: getCurrentTimeInterval
Description: 得到给定时间所处的时间段
InputParameter:
	t: 给定时间
Return: 时间段
*************************************************/
func getCurrentTimeInterval(t time.Time) TimeInterval {
	h, _ := time.ParseDuration("1h")
	begintime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	endtime := begintime.Add(h)
	return TimeInterval{begintime, endtime}
}

// 将TimeInterval划分成1h间隔
func splitTimeInterval(timeinterval TimeInterval) []TimeInterval {
	h, _ := time.ParseDuration("1h")
	rtntimeintervals := []TimeInterval{}
	for btime := timeinterval.Begintime; btime != timeinterval.Endtime; btime = btime.Add(h) {
		etime := btime.Add(h)
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
	mon, _ := time.ParseDuration("24h")
	nowtimeinterval := getCurrentTimeInterval(time.Now())
	endday := nowtimeinterval.Begintime
	for i := 0; i < days; i++ {
		endday = endday.Add(mon)
	}
	nowtimeinterval.Endtime = time.Date(endday.Year(), endday.Month(), endday.Day(), 0, 0, 0, 0, endday.Location())
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
func (t *TimeInterval) AddOneHour() {
	h, _ := time.ParseDuration("1h")
	t.Add(h)
}

// Add 时间相加
func (t *TimeInterval) Add(d time.Duration) {
	t.Begintime = t.Begintime.Add(d)
	t.Endtime = t.Endtime.Add(d)
}

// Equal TimeInterval相等比较
func (t1 *TimeInterval) Equal(t2 TimeInterval) bool {
	return t1.Begintime.Equal(t2.Begintime) &&
		t2.Endtime.Equal(t2.Endtime)
}
