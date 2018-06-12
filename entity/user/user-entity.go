/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user的数据层，保存seat的基本信息，并通过该类型与数据库进行交互
Date: 2018年5月4日 星期五 上午10:53
****************************************************************************/

package user

// Item 用户信息
type Item struct {
	// 用户ID，是唯一主键
	ID string `json:"openID,omitempty" bson:"_id"`
	// 学生netID
	NetID string `json:"netID,omitempty"`
	// hash过的密码
	Hashpassword string `json:"-"`
	// 学生所在学校
	School string `json:"school,omitempty"`
	// 学生被警告次数
	Violation int `json:"violation,omitempty"`
}

// newItem 新建一个UserItem，并返回指针
func newItem(ID string, netID string, hashpassword string,
	school string) *Item {
	newUserItem := new(Item)
	newUserItem.ID = ID
	newUserItem.NetID = netID
	newUserItem.Hashpassword = hashpassword
	newUserItem.School = school
	newUserItem.Violation = 0
	return newUserItem
}

func (item *Item) IsPunished() bool {
	return item.Violation >= 3
}
