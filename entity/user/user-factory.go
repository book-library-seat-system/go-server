/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 工厂包，用于生产对象
Date: 2018年7月8日 星期日 下午1:57
****************************************************************************/

package user

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
