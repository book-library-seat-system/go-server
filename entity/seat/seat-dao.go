/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: seat的dao层，使用数据库层接口，为上层逻辑层提供接口
	PS：该文件所有错误全都panic抛出，不进行错误返回
Date: 2018年5月4日 星期五 上午10:17
****************************************************************************/

package seat

import (
	"errors"
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/book-library-seat-system/go-server/entity/mutexmanager"
	"github.com/book-library-seat-system/go-server/mgdb"
	. "github.com/book-library-seat-system/go-server/util"
)

// seat 使用同一个数据库，但是使用不同的表，每一个表含有一个读写锁
var database *mgo.Database
var locks *mutexmanager.RWMutexManager

// TItemsAtomicService 一个空类型
type TItemsAtomicService struct{}

// service 空类型的指针，使用函数
var service = TItemsAtomicService{}

func init() {
	database = mgdb.Mydb.DB("seat")
	locks = mutexmanager.New()
	service.Insert(newSTItem("testsunyetsenuniversity", 1080))
	fmt.Println("Seat database init!")
}

/*************************************************
Function: Insert
Description: 保存TItems信息到数据库，其中，不会报错
InputParameter:
	titems: 要保存的TItems的信息
Return: none
*************************************************/
func (*TItemsAtomicService) Insert(stitem *STItem) {
	locks.AddLock(stitem.School)
	c := database.C(stitem.School)
	// 加锁解锁
	locks.WLock(stitem.School)
	defer locks.WUnlock(stitem.School)
	for _, titem := range stitem.Titems {
		err := c.Insert(titem)
		// 如果插入重复，忽略此错误，否则抛出
		if err != nil && err.Error()[:6] != "E11000" {
			CheckNewErr(err, "101|数据库座位信息插入出现错误")
		}
	}

}

/*************************************************
Function: FindBySchool
Description: 通过主键School查询数据
InputParameter:
	school: 主键
Return: 找到的对应TItems，如果没有为nil
*************************************************/
func (*TItemsAtomicService) FindBySchool(school string) []TItem {
	c := database.C(school)
	titems := []TItem{}
	locks.RLock(school)
	err := c.Find(nil).All(&titems)
	locks.RUnlock(school)
	CheckNewErr(err, "103|数据库座位信息查找出现错误")
	return titems
}

/*************************************************
Function: FindBySchoolAndTimeInterval
Description: 通过两个主键查询数据
InputParameter:
	school: 主键1
	timeinterval: 主键2
Return: 查找到的座位信息，如果未找到报错
*************************************************/
func (this *TItemsAtomicService) FindBySchoolAndTimeInterval(school string, timeinterval TimeInterval) []Item {
	c := database.C(school)
	titem := TItem{}
	locks.RLock(school)
	err := c.Find(bson.M{"_id": timeinterval}).One(&titem)
	locks.RUnlock(school)
	CheckNewErr(err, "103|数据库座位信息查找出现错误")
	return titem.Items
}

/*************************************************
Function: FindBySchoolAndStudentID
Description: 通过学生id查找关于该学生的座位信息
InputParameter:
	school: 主键
	studentid: 学生id
	seatinfo: 座位预约信息
Return: 查找到的座位信息，如果未找到报错
*************************************************/
func (this *TItemsAtomicService) FindBySchoolAndStudentID(school string, studentid string) []SeatInfo {
	c := database.C(school)
	titems := []TItem{}
	locks.RLock(school)
	err := c.Find(bson.M{"items.studentid": studentid}).Select(bson.M{"items.$": 1}).All(&titems)
	locks.RUnlock(school)
	fmt.Println("Dao:", titems)
	CheckNewErr(err, "103|数据库座位信息查找出现错误")
	sis := make([]SeatInfo, len(titems))
	for i, titem := range titems {
		fmt.Println(i, titem)
		sis[i] = *newSeatInfo(titem.Timeinterval, titem.Items[0])
	}
	fmt.Println(sis)
	return sis
}

/*************************************************
Function: FindOneSeat
Description: 在数据库中寻找一个座位
InputParameter:
	school: 主键1
	timeinterval: 主键2
	seatid: 主键3
Return: 查找的单个座位信息，如果未找到报错
*************************************************/
func (this *TItemsAtomicService) FindOneSeat(school string, timeinterval TimeInterval, seatid int) Item {
	c := database.C(school)
	item := struct {
		Items []Item `json:"items"`
	}{}
	locks.RLock(school)
	err := c.Find(bson.M{
		"_id":   timeinterval,
		"items": bson.M{"$elemMatch": bson.M{"seatid": seatid}},
	}).Select(bson.M{"items.$": 1}).One(&item)
	locks.RUnlock(school)
	if err != nil || len(item.Items) != 1 {
		CheckErr(errors.New("105|不存在该座位"))
	}
	return item.Items[0]
}

/*************************************************
Function: UpdateAllSeat
Description: 通过两个主键更新多个座位信息
InputParameter:
	school: 主键1
	timeinterval: 主键2
	seats: 要更改的座位信息
Return: none
*************************************************/
func (*TItemsAtomicService) UpdateAllSeat(
	school string,
	timeinterval TimeInterval,
	seats []Item) {
	c := database.C(school)
	locks.WLock(school)
	err := c.Update(
		bson.M{"_id": timeinterval},
		bson.M{"$set": bson.M{"items": seats}},
	)
	locks.WUnlock(school)
	CheckNewErr(err, "102|数据库座位信息更新出现错误")
}

/*************************************************
Function: UpdateOneSeat
Description: 通过两个主键更新单个座位信息
InputParameter:
	school: 主键1
	timeinterval: 主键2
	seat: 要更改的座位信息
Return: none
*************************************************/
func (*TItemsAtomicService) UpdateOneSeat(
	school string,
	timeinterval TimeInterval,
	seat Item) {
	c := database.C(school)
	locks.WLock(school)
	err := c.Update(
		bson.M{
			"_id":   timeinterval,
			"items": bson.M{"$elemMatch": bson.M{"seatid": seat.SeatID}},
		},
		bson.M{
			"$set": bson.M{"items.$.seatinfo": seat.Seatinfo, "items.$.studentid": seat.StudentID},
		},
	)
	locks.WUnlock(school)
	CheckNewErr(err, "102|数据库座位信息更新出现错误")
}

/*************************************************
Function: DeleteBySchool
Description: 通过主键School删除数据
InputParameter:
	school: 主键1
Return: none
*************************************************/
func (*TItemsAtomicService) DeleteBySchool(school string) {
	locks.WLock(school)
	err := database.C(school).DropCollection()
	locks.WUnlock(school)
	CheckNewErr(err, "104|数据库座位信息删除出现错误")
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
	c := database.C(school)
	locks.WLock(school)
	err := c.Remove(bson.M{"_id": timeinterval})
	locks.WUnlock(school)
	CheckNewErr(err, "104|数据库座位信息删除出现错误")
}

/*************************************************
Function: DeleteOldTimeInterval
Description: 删除某个时间段之前的所有信息
InputParameter:
	school: 主键1
	timeinterval: 时间段
Return: none
*************************************************/
func (*TItemsAtomicService) DeleteOldTimeInterval(school string, timeinterval TimeInterval) {
	c := database.C(school)
	locks.WLock(school)
	err := c.Remove(bson.M{"_id": bson.M{"$lt": timeinterval}})
	locks.WUnlock(school)
	CheckNewErr(err, "104|数据库座位信息删除出现错误")
}
