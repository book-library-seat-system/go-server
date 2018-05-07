/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的数据层，保存seat的基本信息，并通过该类型与数据库进行交互
Date: 2018年5月4日 星期五 上午10:23
****************************************************************************/

package seat

import "time"

// TimeInterval 时间间隔
type TimeInterval struct {
	// 开始时间
	Begintime time.Time `json:"begintime"`
	// 结束时间
	Endtime time.Time `json:"endtime"`
}

// Item 座位信息
type Item struct {
	// 座位状态
	Seatinfo int `json:"seatinfo"`
	// 外键，预约学生的ID
	StudentID string `json:"studentID"`
}

// TItems School为主键，以TimeInterval作为key值，Item数组作为value值
type TItems struct {
	// 所属学校，主键
	School string
	// 该学校所有可用时间段和座位信息
	Items map[TimeInterval]([]Item)
}

// Equal TimeInterval相等比较
func (t1 TimeInterval) Equal(t2 TimeInterval) bool {
	return t1.Begintime.Equal(t2.Begintime) &&
		t2.Endtime.Equal(t2.Endtime)
}

// //Meeting 会议数据实体
// type Meeting struct {
// 	//会议主题
// 	Title string `xorm:"varchar(64) pk 'title'"`
// 	//会议发起者
// 	Host string `xorm:"varchar(64) notnull 'host'"`
// 	//会议参与者
// 	Participator string `xorm:"varchar(255) 'participators'"`
// 	//开始时间
// 	StartTime time.Time `xorm:"'startime'"`
// 	//结束时间
// 	EndTime time.Time `xorm:"'endTime'"`
// }

// //TableName .
// func (Meeting) TableName() string {
// 	return "meetinginformation"
// }
