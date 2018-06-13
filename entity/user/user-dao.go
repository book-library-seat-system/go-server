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
	"sync"

	"github.com/book-library-seat-system/go-server/mgdb"
	. "github.com/book-library-seat-system/go-server/util"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// user 使用同一个数据库，使用相同的表，表含有一个读写锁
var collector *mgo.Collection
var lock *sync.RWMutex

func init() {
	collector = mgdb.Mydb.DB("user").C("student")
	lock = new(sync.RWMutex)
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
func (*ItemAtomicService) Insert(student *Item) {
	lock.Lock()
	err := collector.Insert(student)
	lock.Unlock()
	CheckNewErr(err, "1|数据库学生信息插入出现错误")
}

/*************************************************
Function: Update
Description: 更新旧项
InputParameter:
	student: 要更新的student
Return: none
*************************************************/
func (*ItemAtomicService) Update(student *Item) {
	lock.Lock()
	err := collector.Update(
		bson.M{"_id": student.ID},
		bson.M{"$set": student},
	)
	lock.Unlock()
	CheckNewErr(err, "2|数据库学生信息更新出现错误")
}

/*************************************************
Function: FindByID
Description: 通过主键ID查询数据
InputParameter:
	ID: 学生的ID
Return: 查询到的学生结果，包含所有的字段
*************************************************/
func (*ItemAtomicService) FindByID(ID string) *Item {
	item := &Item{}
	lock.RLock()
	err := collector.Find(bson.M{"_id": ID}).One(item)
	lock.RUnlock()
	CheckNewErr(err, "3|数据库学生信息查找出现错误")
	return item
}

/*************************************************
Function: FindSchoolByID
Description: 通过主键ID查询数据
InputParameter:
	ID: 学生的ID
Return: 查询到的学生学校结果
*************************************************/
func (*ItemAtomicService) FindSchoolByID(ID string) string {
	item := &Item{}
	lock.RLock()
	err := collector.Find(bson.M{"_id": ID}).Select(bson.M{"school": 1}).One(item)
	lock.RUnlock()
	CheckNewErr(err, "3|数据库学生信息查找出现错误")
	return item.School
}

/*************************************************
Function: DeleteByID
Description: 通过主键ID删除数据，是真正的删除
InputParameter:
	ID: 学生的ID
Return: none
*************************************************/
func (*ItemAtomicService) DeleteByID(ID string) {
	lock.Lock()
	err := collector.Remove(bson.M{"_id": ID})
	lock.Unlock()
	CheckNewErr(err, "4|数据库学生信息删除出现错误")
}
