package seat

import (
	"errors"
	"testing"

	. "github.com/book-library-seat-system/go-server/util"
)

var sysustitem *STItem
var whstitem *STItem

func init() {
	// 新生成中山大学服务和武汉大学服务
	sysustitem = newSTItem("testschoolsysu", 10)
	//fmt.Println(sysustitem)
	whstitem = newSTItem("testschoolwu", 15)
}

// 判断是否相同，不相同报错
func equalTItems(orititems []TItem, rtntitems []TItem) {
	// 判断TItem数组长度
	if len(orititems) != len(rtntitems) {
		panic(errors.New("TItems: []TItem hasn't same length!"))
	}

	// 递归循环判断
	for i := 0; i < len(orititems); i++ {
		oriitems := orititems[i].Items
		rtnitems := rtntitems[i].Items

		// 判断每个TItem中的时间段长度
		if !orititems[i].Timeinterval.Equal(rtntitems[i].Timeinterval) {
			panic(errors.New("TItem: TItem hasn't same timeinterval!"))
		}

		// 判断Item数组是否相等
		equalItems(oriitems, rtnitems)
	}
}

func equalItems(oriitems []Item, rtnitems []Item) {
	// 判断每个TItem中的Item数组长度
	if len(oriitems) != len(rtnitems) {
		panic(errors.New("Items: []Item hasn't same length!"))
	}

	// 	判断座位是否匹配
	for j := 0; j < len(oriitems); j++ {
		equalItem(oriitems[j], rtnitems[j])
	}
}

func equalItem(oriitem Item, rtnitem Item) {
	if oriitem.SeatID != rtnitem.SeatID ||
		oriitem.Seatinfo != rtnitem.Seatinfo ||
		oriitem.StudentID != rtnitem.StudentID {
		panic(errors.New("Item: Item hasn't same info!"))
	}
}

func TestInsert(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 修改其中一项数据
	sysustitem.Titems[5].Items[5].Seatinfo = 1
	sysustitem.Titems[5].Items[5].StudentID = "15331111"

	whstitem.Titems[3].Items[10].Seatinfo = 2
	whstitem.Titems[3].Items[10].StudentID = "15331111"

	service.Insert(sysustitem)
	service.Insert(whstitem)

	// 从数据库读取数据并进行比较
	c := database.C(sysustitem.School)
	rtntitems := []TItem{}
	err := c.Find(nil).All(&rtntitems)
	CheckErr(err)
	equalTItems(sysustitem.Titems, rtntitems)

	c = database.C(whstitem.School)
	err = c.Find(nil).All(&rtntitems)
	CheckErr(err)
	equalTItems(whstitem.Titems, rtntitems)
}

func TestFindBySchool(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 一个一个测试
	titems := service.FindBySchool(sysustitem.School)
	equalTItems(sysustitem.Titems, titems)
	titems = service.FindBySchool(whstitem.School)
	equalTItems(whstitem.Titems, titems)
}

func TestFindBySchoolAndTimeInterval(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 测试两个特例
	items := service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[5].Timeinterval)
	equalItems(sysustitem.Titems[5].Items, items)
	items = service.FindBySchoolAndTimeInterval(whstitem.School, whstitem.Titems[3].Timeinterval)
	equalItems(whstitem.Titems[3].Items, items)
}

func TestFindOneSeat(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	item := service.FindOneSeat(sysustitem.School, sysustitem.Titems[5].Timeinterval, 5)
	equalItem(sysustitem.Titems[5].Items[5], item)
	item = service.FindOneSeat(whstitem.School, whstitem.Titems[3].Timeinterval, 10)
	equalItem(whstitem.Titems[3].Items[10], item)
}

func TestUpdateAllSeat(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 更新两个座位信息，并插入
	sysustitem.Titems[4].Items[4].Seatinfo = 2
	sysustitem.Titems[4].Items[4].StudentID = "11112222"
	sysustitem.Titems[4].Items[6].Seatinfo = 1
	sysustitem.Titems[4].Items[6].StudentID = "22222222"
	service.UpdateAllSeat(sysustitem.School, sysustitem.Titems[4].Timeinterval, sysustitem.Titems[4].Items)
	items := service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[4].Timeinterval)
	equalItems(sysustitem.Titems[4].Items, items)
}

func TestUpdateOneSeat(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 更新一个座位信息，并插入
	sysustitem.Titems[3].Items[3].Seatinfo = 1
	sysustitem.Titems[3].Items[3].StudentID = "11111111"
	service.UpdateOneSeat(sysustitem.School, sysustitem.Titems[3].Timeinterval, sysustitem.Titems[3].Items[3])
	items := service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[3].Timeinterval)
	equalItems(sysustitem.Titems[3].Items, items)
}

func TestDeleteBySchoolAndTimeInterval(t *testing.T) {
	defer func() {
		if err := recover(); err != nil && err.(error).Error()[:3] != "103" {
			t.Error(err)
		}
	}()

	// 删除某个时间段再进行查找
	service.DeleteBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[0].Timeinterval)
	service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[0].Timeinterval)
}

func TestDeleteOldTimeInterval(t *testing.T) {
	defer func() {
		if err := recover(); err != nil && err.(error).Error()[:3] != "103" {
			t.Error(err)
		}
	}()

	// 删除某个时间段再进行查找
	service.DeleteOldTimeInterval(sysustitem.School, sysustitem.Titems[5].Timeinterval)
	service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[4].Timeinterval)
}

func TestDeleteBySchool(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 删除所有内容
	service.DeleteBySchool(sysustitem.School)
	service.DeleteBySchool(whstitem.School)
	names, err := database.CollectionNames()
	CheckErr(err)
	if len(names) != 0 {
		t.Error(errors.New("DeleteBySchool error!"))
	}
}
