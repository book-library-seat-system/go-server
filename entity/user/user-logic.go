package user

import (
	"errors"
	"path/filepath"

	"github.com/book-library-seat-system/go-server/entity/mylog"
)

var userItemsFilePath = "src/github.com/book-library-seat-system/go-server/orm/UserItems.json"
var currentUserFilePath = "src/github.com/book-library-seat-system/go-server/orm/Current.txt"

func init() {
	// 初始化
	userItemsFilePath = filepath.Join(*mylog.GetGOPATH(), userItemsFilePath)
	currentUserFilePath = filepath.Join(*mylog.GetGOPATH(), currentUserFilePath)
}

// hasUser : 判断user是否存在
func hasUser(ID string) bool {
	return service.FindByID(ID) != nil
}

// RegisterUser : 注册用户，如果用户名一样，则panic错误
func RegisterUser(ID string, name string, password string,
	email string, school string) {
	if hasUser(ID) {
		checkErr(errors.New("1|ERROR:The user has registered"))
	}

	pitem := newItem(ID, name, password, email, school)
	service.Save(pitem)

	// mylog.AddLog("", "RegisterUser", "", pitem.String())
}

// LoginUser : 登录用户
// 如果用户名不存在，则返回err1
// 或者用户名密码不正确，则返回err2
func LoginUser(ID string, password string) *Item {
	pitem := service.FindByID(ID)
	// 账号错误
	if pitem == nil {
		checkErr(errors.New("2|ERROR:The user's name not exists"))
	}

	// 密码错误
	if pitem.Hashpassword != password {
		checkErr(errors.New("3|ERROR:The user's password is wrong"))
	}

	// 成功登录
	// mylog.AddLog(name, "LoginUser", "", "")
	return pitem
}

// LogoutUser : 登出用户，如果当前没有用户登录，则返回err
func LogoutUser(ID string) {
	if !hasUser(ID) {
		checkErr(errors.New("4|ERROR:No login user"))
	}

	// mylog.AddLog(ID, "LogoutUser", "", "")
}

// DeleteUser : 删除当前登录用户，删除后当前登录用户置为nil
// 如果当前没有用户登录，返回err
// func DeleteUser(loginname string) {
// 	if !IsLogin(loginname) {
// 		checkErr(errors.New("ERROR:No registered user"))
// 	}

// 	service.DeleteByName(loginname)
// 	mylog.AddLog(loginname, "DeleteUser", "", "")
// }

// to string
func (u Item) String() string {
	return "{Name:" + u.Name + "  Email:" + u.Email + "  Phone:" + u.School + "}"
}
