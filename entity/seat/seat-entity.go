/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的数据层，保存seat的基本信息，并通过该类型与数据库进行交互
Date: 2018年5月4日 星期五 上午10:23
****************************************************************************/

package seat

import (
	"errors"
	"strconv"
	"time"

	. "github.com/book-library-seat-system/go-server/util"
	"github.com/vaughan0/go-ini"
)

// Item 座位信息
type Item struct {
	// 座位唯一ID
	SeatID int `json:"seatID"`
	// 座位状态
	Seatinfo int `json:"seatinfo"`
	// 外键，预约学生的ID
	StudentID string `json:"studentID"`
}

// TimeInterval 时间间隔
type TimeInterval struct {
	// 开始时间
	Begintime time.Time `json:"begintime"`
	// 结束时间
	Endtime time.Time `json:"endtime"`
}

// TItem 包含时间间隔和该时间间隔相应的座位信息
type TItem struct {
	Timeinterval TimeInterval `json:"timeinterval"`
	// 座位信息
	Items []Item `json:"item"`
}

// STItem 包含学校姓名和该学校所有的座位信息
type STItem struct {
	// 所属学校，主键
	School string `bson:"_id" json:"school"`
	// 该学校所有可用时间段和座位信息
	Titems []TItem `json:"titems"`
}

// newItems 生成一个Item数组，id从0开始
func newItems(seatnumber int) []Item {
	items := make([]Item, seatnumber)
	for i := 0; i < len(items); i++ {
		items[i].SeatID = i
		items[i].Seatinfo = 0
		items[i].StudentID = ""
	}
	return items
}

// newTItems 生成一个TItem数组，timeinterval从当前时间段开始，数组数量从配置文件读取
func newTItems(seatnumber int) []TItem {
	titems := []TItem{}
	timeintervals := currentTimeInterval()
	for i := 0; i < len(timeintervals); i++ {
		titems = append(titems, TItem{
			Timeinterval: timeintervals[i],
			Items:        newItems(seatnumber),
		})
	}
	return titems
}

// newSTItem 生成一个STItem
func newSTItem(school string, seatnumber int) *STItem {
	newtitems := new(STItem)
	newtitems.School = school
	newtitems.Titems = newTItems(seatnumber)
	return newtitems
}

// 通过配置文件，读取有效时间段
func currentTimeInterval() []TimeInterval {
	// 从配置文件中读取段数
	file, err := ini.LoadFile("config.ini")
	CheckErr(err)
	timesstr, ok := file.Get("TimeInterval", "times")
	if !ok {
		panic(errors.New("config.ini haven't \"times\"!"))
	}
	times, err := strconv.Atoi(timesstr)
	CheckErr(err)

	// 生成时间段
	timeinterval := []TimeInterval{}
	for i := 0; i < times; i++ {

	}
	return timeinterval
}

// Equal TimeInterval相等比较
func (t1 TimeInterval) Equal(t2 TimeInterval) bool {
	return t1.Begintime.Equal(t2.Begintime) &&
		t2.Endtime.Equal(t2.Endtime)
}
