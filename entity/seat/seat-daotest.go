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
	sysustitem = newSTItem("sunyetsununiversity", 10)
	whstitem = newSTItem("wuhanuniversity", 15)
}

// 判断是否相同，不相同报错
func equalTItem(orititems []TItem, rtntitems []TItem) {
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
		equalItem(oriitems, rtnitems)
	}
}

func equalItem(oriitems []Item, rtnitems []Item) {
	// 判断每个TItem中的Item数组长度
	if len(oriitems) != len(rtnitems) {
		panic(errors.New("Items: []Item hasn't same length!"))
	}

	// 	判断座位是否匹配
	for j := 0; j < len(oriitems); j++ {
		if oriitems[j].SeatID != rtnitems[j].SeatID ||
			oriitems[j].Seatinfo != rtnitems[j].Seatinfo ||
			oriitems[j].StudentID != rtnitems[j].StudentID {
			panic(errors.New("Item: Item hasn't same info!"))
		}
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
	equalTItem(sysustitem.Titems, rtntitems)

	c = database.C(whstitem.School)
	err = c.Find(nil).All(&rtntitems)
	CheckErr(err)
	equalTItem(whstitem.Titems, rtntitems)
}

// func TestFindAll(t *testing.T) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			t.Error(err)
// 		}
// 	}()

// 	// 测试两者等不等价
// 	stitems := service.FindAll()
// 	//fmt.Println(len(stitems))
// 	if stitems[0].School == whstitem.School {
// 		equalTItem(whstitem.Titems, stitems[0].Titems)
// 		equalTItem(sysustitem.Titems, stitems[1].Titems)
// 	} else if stitems[0].School == sysustitem.School {
// 		equalTItem(sysustitem.Titems, stitems[0].Titems)
// 		equalTItem(whstitem.Titems, stitems[1].Titems)
// 	} else {
// 		t.Error("FindAll error!")
// 	}
// }

func TestFindBySchool(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 一个一个测试
	titems := service.FindBySchool(sysustitem.School)
	equalTItem(sysustitem.Titems, titems)
	titems = service.FindBySchool(whstitem.School)
	equalTItem(whstitem.Titems, titems)
}

func TestFindBySchoolAndTimeInterval(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	// 测试两个特例
	items := service.FindBySchoolAndTimeInterval(sysustitem.School, sysustitem.Titems[5].Timeinterval)
	equalItem(sysustitem.Titems[5].Items, items)
	items = service.FindBySchoolAndTimeInterval(whstitem.School, whstitem.Titems[3].Timeinterval)
	equalItem(whstitem.Titems[3].Items, items)
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
	equalItem(sysustitem.Titems[4].Items, items)
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
	equalItem(sysustitem.Titems[3].Items, items)
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
