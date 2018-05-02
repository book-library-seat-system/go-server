package user

import (
	"github.com/book-library-seat-system/go-server/entity/mylog"
	"github.com/book-library-seat-system/go-server/orm"
)

// ItemAtomicService 一个空类型
type ItemAtomicService struct{}

// service 空类型的指针，使用函数
var service = ItemAtomicService{}

func init() {
	err := orm.Mydb.Sync2(new(Item))
	checkErr(err)
}

// Save 保存
func (*ItemAtomicService) Save(u *Item) error {
	_, err := orm.Mydb.Table("user").Insert(u)
	checkErr(err)
	return err
}

// FindAll 找到所有Item
func (*ItemAtomicService) FindAll() []Item {
	as := []Item{}
	err := orm.Mydb.Table("user").Desc("ID").Find(&as)
	checkErr(err)
	return as
}

// FindByName 通过主键ID查询数据
func (*ItemAtomicService) FindByID(ID string) *Item {
	a := &Item{}
	_, err := orm.Mydb.Table("user").Id(ID).Get(a)
	checkErr(err)
	return a
}

// DeleteByName 通过主键ID删除数据
func (*ItemAtomicService) DeleteByID(ID string) {
	// 软删除
	_, err := orm.Mydb.Table("user").Id(ID).Delete(&Item{})
	checkErr(err)

	// 真正删除
	_, err = orm.Mydb.Table("user").Id(ID).Unscoped().Delete(&Item{})
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		mylog.AddErr(err)
		panic(err)
	}
}
