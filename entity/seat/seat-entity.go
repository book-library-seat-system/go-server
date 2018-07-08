/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的数据层，保存seat的基本信息，并通过该类型与数据库进行交互
Date: 2018年5月4日 星期五 上午10:23
****************************************************************************/

package seat

import (
	"time"
)

const (
	// UnBook 座位未预约状态
	UnBook = 0
	// Book 座位已被预约状态
	Book = 1
	// Signin 座位被签到状态
	Signin = 2
	// Signout 座位已被签退状态
	Signout = 3
	// BookAndUnSignin 座位被预约却没被签到，惩罚状态
	BookAndUnSignin = 4
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
	// 时间信息
	Timeinterval TimeInterval `json:"timeinterval" bson:"_id"`
	// 座位信息
	Items []Item `json:"item"`
}

// SeatInfo 包含时间段信息和座位ID，用于显示
type SeatInfo struct {
	// 时间信息
	TimeInterval
	// 座位唯一ID
	SeatID int `json:"seatID"`
	// 座位状态
	Seatinfo int `json:"seatinfo"`
}

// STItem 包含学校姓名和该学校所有的座位信息
type STItem struct {
	// 所属学校，主键
	School string `json:"school"`
	// 该学校所有可用时间段和座位信息
	Titems []TItem `json:"titems"`
}
