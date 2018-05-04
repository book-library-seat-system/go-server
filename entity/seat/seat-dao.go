/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的dao层，使用数据库层接口，为上层逻辑层提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:17
****************************************************************************/

package seat

import (
	"github.com/book-library-seat-system/go-server/orm"
	. "github.com/book-library-seat-system/go-server/util"
)

func init() {
	// err := orm.Mydb.Sync2(new(TItems))
	// CheckErr(err)
	// fmt.Println("Meeting database init!")
}

// TItemsAtomicService 一个空类型
type TItemsAtomicService struct{}

// service 空类型的指针，使用函数
var service = TItemsAtomicService{}

/*************************************************
Function: Insert
Description: 保存TItems信息到数据库
InputParameter:
	titems: 要保存的TItems的信息
Return: none
*************************************************/
func (*TItemsAtomicService) Insert(titems *TItems) {
	_, err := orm.Mydb.Table("tseats").Insert(titems)
	CheckErr(err)
}

/*************************************************
Function: FindAll
Description: 找到所有Item
InputParameter: none
Return: 找到的所有TItems列表
*************************************************/
func (*TItemsAtomicService) FindAll() []TItems {
	as := []TItems{}
	err := orm.Mydb.Table("tseats").Desc("TimeInterval").Find(&as)
	CheckErr(err)
	return as
}

/*************************************************
Function: FindBySchool
Description: 通过主键School查询数据
InputParameter:
	school: 主键
Return: 找到的对应TItems，如果没有为nil
*************************************************/
func (*TItemsAtomicService) FindBySchool(school string) TItems {
	a := TItems{}
	_, err := orm.Mydb.Table("tseats").Id(school).Get(a)
	CheckErr(err)
	return a
}

/*************************************************
Function: FindBySchoolAndTimeInterval
Description: 通过两个主键查询数据
InputParameter:
	school: 主键1
	timeinterval: 主键2
Return: 查找到的座位信息，如果不存在返回nil
*************************************************/
func (*TItemsAtomicService) FindBySchoolAndTimeInterval(school string, timeinterval TimeInterval) []Item {
	return nil
}

/*************************************************
Function: DeleteBySchool
Description: 通过主键School删除数据
InputParameter:
	school: 主键1
Return: none
*************************************************/
//
func (*TItemsAtomicService) DeleteBySchool(school string) {
	// // 软删除
	// _, err := orm.Mydb.Table("tseats").Id(timeinterval).Delete(&Item{})
	// CheckErr(err)

	// // 真正删除
	// _, err = orm.Mydb.Table("tseats").Id(timeinterval).Unscoped().Delete(&Item{})
	// CheckErr(err)
}

/*************************************************
Function: DeleteBySchoolAndTimeInterval
Description: 通过两个主键删除数据（删除过时信息）
InputParameter:
	school: 主键1
	timeinterval: 主键2
Return: none
*************************************************/
func (*TItemsAtomicService) DeleteBySchoolAndTimeInterval(school string, timeinterval TimeInterval) {
	// // 软删除
	// _, err := orm.Mydb.Table("tseats").Id(timeinterval).Delete(&Item{})
	// CheckErr(err)

	// // 真正删除
	// _, err = orm.Mydb.Table("tseats").Id(timeinterval).Unscoped().Delete(&Item{})
	// CheckErr(err)
}

/*************************************************
Function: UpdateBySchoolAndTimeInterval
Description: 通过两个主键更新座位信息
InputParameter:
	school: 主键1
	timeinterval: 主键2
	seats: 要更改的座位信息
Return: none
*************************************************/
func (*TItemsAtomicService) UpdateBySchoolAndTimeInterval(
	school string,
	timeinterval TimeInterval,
	seats []Item) {

}
