package user

import (
	"fmt"
	"testing"

	"github.com/book-library-seat-system/go-server/mgdb"
	. "github.com/book-library-seat-system/go-server/util"
)

var hza *Item
var lbb *Item
var hmy *Item

func init() {
	hza = newItem("12345678", "huziang", MD5Hash("111"), "111@qq.com", "Sun-Yet Sun University")
	lbb = newItem("11112222", "linbinbin", MD5Hash("222"), "222@qq.com", "Sun-Yet Sun University")
	hmy = newItem("33334444", "huangminyi", MD5Hash("333"), "333@qq.com", "Sun-Yet Sun University")

	c := mgdb.Mydb.DB("user").C("student")
	c.RemoveAll(nil)
}

func TestSave(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 插入三个节点
	service.Save(hza)
	service.Save(lbb)
	service.Save(hmy)

	c := mgdb.Mydb.DB("user").C("student")
	users := &[]Item{}
	c.Find(nil).All(users)
	if len(*users) != 3 {
		t.Error("Save error!")
	}
}

func TestFindByID(t *testing.T) {
	equal := func(item1 *Item, item2 *Item) bool {
		return item1.ID == item2.ID &&
			item1.Name == item2.Name &&
			item1.Hashpassword == item2.Hashpassword &&
			item1.Email == item2.Email &&
			item1.School == item2.School &&
			item1.Violation == item2.Violation
	}

	// 查看结果是否相同
	user := service.FindByID(hza.ID)
	if !equal(user, hza) {
		t.Error("Find error!")
	}
	user = service.FindByID(lbb.ID)
	if !equal(user, lbb) {
		t.Error("Find error!")
	}
	user = service.FindByID(hmy.ID)
	if !equal(user, hmy) {
		t.Error("Find error!")
	}
}

func TestUpdate(t *testing.T) {
	// 修改一个属性，判断数据库有没有修改
	hza.Email = "qqq@qq.com"
	service.Update(hza)
	user := service.FindByID(hza.ID)
	if user.Email != "qqq@qq.com" {
		t.Error("Update error!")
	}
}

func TestDeleteByID(t *testing.T) {
	// 删除所有信息
	//service.DeleteByID(hza.ID)
	service.DeleteByID(lbb.ID)
	service.DeleteByID(hmy.ID)

	c := mgdb.Mydb.DB("user").C("student")
	users := &[]Item{}
	c.Find(nil).All(users)
	if len(*users) != 1 {
		t.Error("Save error!")
	}
}

func TestShow(t *testing.T) {
	c := mgdb.Mydb.DB("user").C("student")
	users := &[]Item{}
	c.Find(nil).All(users)
	for i := 0; i < len(*users); i++ {
		fmt.Println((*users)[i])
	}
}
