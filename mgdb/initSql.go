/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: mongodb数据库初始化部分，和dao层对接
Date: 2018年5月4日 星期五 下午1:15
****************************************************************************/

package mgdb

import (
	. "github.com/book-library-seat-system/go-server/util"
	"labix.org/v2/mgo"
)

// Mydb 数据库指针
var Mydb *mgo.Session

// 生成数据库，对数据库进行链接
func init() {
	// 链接mongodb数据库
	session, err := mgo.Dial("")
	CheckErr(err)
	session.SetMode(mgo.Monotonic, true)

	Mydb = session
}
