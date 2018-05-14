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
Function: GetStudent
Description: 返回该学生的所有信息
InputParameter:
	ID: 学生ID
Return: 该学生的所有信息
*************************************************/
func GetStudent(ID string) *Item {
	return service.FindByID(ID)
}

/*************************************************
Function: RegisterStudent
Description: 注册用户
InputParameter:
	ID: 学生ID
	name: 学生姓名
	password: 学生密码（哈希过后的密码）
	email: 学生邮箱
	school: 学生学校
Return: none
*************************************************/
func RegisterStudent(ID string, name string, password string,
	email string, school string) {
	pitem := newItem(ID, name, password, email, school)
	service.Insert(pitem)
}

/*************************************************
Function: UpdateStudent
Description: 更新学生信息
InputParameter:
	ID: 学生ID
	name: 学生姓名
	password: 学生密码（哈希过后的密码）
	email: 学生邮箱
	school: 学生学校
Return: none
*************************************************/
func UpdateStudent(ID string, name string, password string,
	email string, school string) {
	pitem := newItem(ID, name, password, email, school)
	service.Update(pitem)
}

/*************************************************
Function: LoginStudent
Description: 学生登录
InputParameter:
	ID: 学生ID
	password: 学生密码（哈希过后）
Return: 该学生的所有信息
*************************************************/
func LoginStudent(ID string, password string) *Item {
	pitem := GetStudent(ID)
	// 密码错误
	if pitem.Hashpassword != password {
		CheckErr(errors.New("5|学生密码错误"))
	}
	return pitem
}

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
