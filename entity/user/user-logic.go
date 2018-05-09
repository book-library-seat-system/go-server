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
Function: HasStudent
Description: 判断user是否存在
InputParameter:
	ID: 学生ID
Return: 是否存在该学生
*************************************************/
func HasStudent(ID string) bool {
	return service.FindByID(ID) != nil
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
	if HasStudent(ID) {
		CheckErr(errors.New("1|ERROR:The user has registered"))
	}

	pitem := newItem(ID, name, password, email, school)
	service.Insert(pitem)

	// mylog.AddLog("", "RegisterUser", "", pitem.String())
}

/*************************************************
Function: UpdateStudent
Description: 更新学生信息
InputParameter:
	student: 学生更新后的信息
Return: none
*************************************************/
func UpdateStudent(student Item) {
}

/*************************************************
Function: LoginStudent
Description: 学生登录
InputParameter:
	ID: 学生ID
	password: 学生密码（哈希过后）
Return: none
*************************************************/
func LoginStudent(ID string, password string) {
	pitem := GetStudent(ID)
	// 账号错误
	if pitem == nil {
		CheckErr(errors.New("2|ERROR:The user's name not exists"))
	}

	// 密码错误
	if pitem.Hashpassword != password {
		CheckErr(errors.New("3|ERROR:The user's password is wrong"))
	}

	// 成功登录
	// mylog.AddLog(name, "LoginUser", "", "")
}

/*************************************************
Function: DeleteStudent
Description: 通过ID删除用户
InputParameter:
	ID: 学生ID
Return: none
*************************************************/
func DeleteStudent(ID string) {

}

// // LogoutStudent : 登出用户
// func LogoutStudent(ID string) {
// 	if !HasStudent(ID) {
// 		CheckErr(errors.New("4|ERROR:No login user"))
// 	}

// 	// mylog.AddLog(ID, "LogoutUser", "", "")
// }

// to string
// func (u Item) String() string {
// 	return "{Name:" + u.Name + "  Email:" + u.Email + "  Phone:" + u.School + "}"
// }
