/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user的dao层，使用数据库层接口，为上层逻辑层提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:51
****************************************************************************/

package user

import (
	"fmt"

	"github.com/book-library-seat-system/go-server/mgdb"
	. "github.com/book-library-seat-system/go-server/util"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var collector *mgo.Collection

func init() {
	collector = mgdb.Mydb.DB("user").C("student")
	fmt.Println("User database init!")
}

// ItemAtomicService 一个空类型
type ItemAtomicService struct{}

// service 空类型的指针，使用函数
var service = ItemAtomicService{}

/*************************************************
Function: Insert
Description: 插入新项
InputParameter:
	student: 新的student
Return: none
*************************************************/
func (*ItemAtomicService) Save(student *Item) {
	err := collector.Insert(student)
	CheckDBErr(err, "1|数据库学生信息插入出现错误")
}

/*************************************************
Function: Update
Description: 更新旧项
InputParameter:
	student: 要更新的student
Return: none
*************************************************/
func (*ItemAtomicService) Update(student *Item) {
	err := collector.Update(
		bson.M{"_id": student.ID},
		bson.M{"$set": student},
	)
	CheckDBErr(err, "2|数据库学生信息更新出现错误")
}

/*************************************************
Function: FindByName
Description: 通过主键ID查询数据
InputParameter:
	ID: 学生的ID
Return: 查询到的学生结果，包含所有的字段
*************************************************/
func (*ItemAtomicService) FindByID(ID string) *Item {
	item := &Item{}
	err := collector.Find(bson.M{"_id": ID}).One(item)
	CheckDBErr(err, "3|数据库学生信息查找出现错误")
	return item
}

/*************************************************
Function: DeleteByName
Description: 通过主键ID删除数据，是真正的删除
InputParameter:
	ID: 学生的ID
Return: none
*************************************************/
func (*ItemAtomicService) DeleteByID(ID string) {
	err := collector.Remove(bson.M{"_id": ID})
	CheckDBErr(err, "4|数据库学生信息删除出现错误")
}
