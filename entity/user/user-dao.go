/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: user的dao层，使用数据库层接口，为上层逻辑层提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:51
****************************************************************************/

package user

import (
	"github.com/book-library-seat-system/go-server/orm"
	. "github.com/book-library-seat-system/go-server/util"
)

func init() {
	// err := orm.Mydb.Sync2(new(Item))
	// CheckErr(err)
	// fmt.Println("User database init!")
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
	_, err := orm.Mydb.Table("user").Insert(n)
	CheckErr(err)
}

/*************************************************
Function: Update
Description: 更新旧项
InputParameter:
	student: 要更新的student
Return: none
*************************************************/
func (*ItemAtomicService) Update(student *Item) {

}

// FindAll 找到所有Item
// func (*ItemAtomicService) FindAll() []Item {
// 	as := []Item{}
// 	err := orm.Mydb.Table("user").Desc("ID").Find(&as)
// 	CheckErr(err)
// 	return as
// }

/*************************************************
Function: FindByName
Description: 通过主键ID查询数据
InputParameter:
	ID: 学生的ID
Return: 查询到的学生结果，包含所有的字段
*************************************************/
func (*ItemAtomicService) FindByID(ID string) *Item {
	a := &Item{}
	_, err := orm.Mydb.Table("user").Id(ID).Get(a)
	CheckErr(err)
	return a
}

/*************************************************
Function: DeleteByName
Description: 通过主键ID删除数据，是真正的删除
InputParameter:
	ID: 学生的ID
Return: none
*************************************************/
func (*ItemAtomicService) DeleteByID(ID string) {
	// 软删除
	_, err := orm.Mydb.Table("user").Id(ID).Delete(&Item{})
	CheckErr(err)

	// 真正删除
	_, err = orm.Mydb.Table("user").Id(ID).Unscoped().Delete(&Item{})
	CheckErr(err)
}
