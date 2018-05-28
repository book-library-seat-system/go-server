package user

import (
	"testing"

	"github.com/book-library-seat-system/go-server/mgdb"
	. "github.com/book-library-seat-system/go-server/util"
)

var hza *Item
var lbb *Item
var hmy *Item

func init() {
	hza = newItem("12345678", "huziang", MD5Hash("111"), "Sun-Yet Sun University")
	lbb = newItem("11112222", "linbinbin", MD5Hash("222"), "Sun-Yet Sun University")
	hmy = newItem("33334444", "huangminyi", MD5Hash("333"), "Sun-Yet Sun University")
}

func TestSave(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 插入三个节点
	service.Insert(hza)
	service.Insert(lbb)
	service.Insert(hmy)

	c := mgdb.Mydb.DB("user").C("student")
	users := &[]Item{}
	c.Find(nil).All(users)
	if len(*users) != 3 {
		t.Error("Save error!")
	}
}

func TestFindByID(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	equal := func(item1 *Item, item2 *Item) bool {
		return item1.ID == item2.ID &&
			item1.NetID == item2.NetID &&
			item1.Hashpassword == item2.Hashpassword &&
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

func TestFindSchoolByID(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 查看结果是否相同
	school := service.FindSchoolByID(hza.ID)
	if school != hza.School {
		t.Error("Find error!")
	}
	school = service.FindSchoolByID(lbb.ID)
	if school != lbb.School {
		t.Error("Find error!")
	}
	school = service.FindSchoolByID(hmy.ID)
	if school != hmy.School {
		t.Error("Find error!")
	}
}

func TestUpdate(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 修改一个属性，判断数据库有没有修改
	hza.NetID = "hza"
	service.Update(hza)
	user := service.FindByID(hza.ID)
	if user.NetID != "hza" {
		t.Error("Update error!")
	}
}

func TestDeleteByID(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 删除全部
	service.DeleteByID(hza.ID)
	service.DeleteByID(lbb.ID)
	service.DeleteByID(hmy.ID)

	c := mgdb.Mydb.DB("user").C("student")
	users := &[]Item{}
	c.Find(nil).All(users)
	if len(*users) != 0 {
		t.Error("Delete error!")
	}
}
