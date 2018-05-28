/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user的逻辑层，使用dao层接口，为上层控制层（server层）提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:52
****************************************************************************/

package user

import (
	"errors"

	. "github.com/book-library-seat-system/go-server/util"
)

var userItemsFilePath = "src/github.com/book-library-seat-system/go-server/orm/UserItems.json"
var currentUserFilePath = "src/github.com/book-library-seat-system/go-server/orm/Current.txt"

func init() {
	// 初始化
	// userItemsFilePath = filepath.Join(*mylog.GetGOPATH(), userItemsFilePath)
	// currentUserFilePath = filepath.Join(*mylog.GetGOPATH(), currentUserFilePath)
}

/*************************************************
Function: testStudent
Description: 测试给定的用户ID和密码，判断是否有这个人存在
InputParameter:
	ID: 学生ID
	password: 密码
	school: 学校
Return: 学校是否存在学生
*************************************************/
func testStudent(ID string, password string, school string) bool {
	return true
}

/*************************************************
Function: GetStudent
Description: 返回该学生的所有信息
InputParameter:
	openID: 微信OpenID
Return: 该学生的所有信息
*************************************************/
func GetStudent(openID string) *Item {
	return service.FindByID(openID)
}

/*************************************************
Function: GetStudentsSchool
Description: 返回该学生的学校信息
InputParameter:
	openID: 微信OpenID
Return: 该学生的学校信息
*************************************************/
func GetStudentsSchool(openID string) string {
	return service.FindSchoolByID(openID)
}

/*************************************************
Function: RegisterStudent
Description: 注册用户
InputParameter:
	openID: 微信OpenID
	netID: 中山大学netID
	hashpassword: 学生密码（哈希过后的密码）
	school: 学生学校
Return: none
*************************************************/
func RegisterStudent(openID string, netID string, hashpassword string,
	school string) {
	pitem := newItem(openID, netID, hashpassword, school)
	if !testStudent(netID, hashpassword, school) {
		CheckErr(errors.New("200|未定义错误"))
	}
	service.Insert(pitem)
}

/*************************************************
Function: UpdateStudent
Description: 更新学生信息
InputParameter:
	openID: 微信OpenID
	netID: 中山大学netID
	hashpassword: 学生密码（哈希过后的密码）
	school: 学生学校
Return: none
*************************************************/
func UpdateStudent(openID string, netID string, hashpassword string,
	school string) {
	pitem := newItem(openID, netID, hashpassword, school)
	service.Update(pitem)
}

/*************************************************
Function: LoginStudent
Description: 学生登录
InputParameter:
	openID: 微信OpenID
Return: 该学生的所有信息
*************************************************/
// func LoginStudent(ID string) *Item {
// 	pitem := GetStudent(ID)
// 	// 密码错误
// 	// if pitem.Hashpassword != password {
// 	// 	CheckErr(errors.New("5|学生密码错误"))
// 	// }
// 	return pitem
// }

/*************************************************
Function: DeleteStudent
Description: 通过ID删除用户
InputParameter:
	ID: 学生ID
Return: none
*************************************************/
func DeleteStudent(ID string) {
	service.DeleteByID(ID)
}
