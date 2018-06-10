/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 用于初始化数据库和触发器所使用
Date: 2018年6月10日 星期日 下午3:34
****************************************************************************/

package seat

import (
	"fmt"
	"time"

	"github.com/book-library-seat-system/go-server/entity/timetrigger"
	"github.com/book-library-seat-system/go-server/entity/user"
	"github.com/book-library-seat-system/go-server/mgdb"
)

func init() {
	initDao()
	addSignoutTrigger("testsunyetsununiversity")
	addSeatUpdateTrigger("testsunyetsununiversity")
	//addTestTrigger()
}

// initDao 初始化数据库
func initDao() {
	database = mgdb.Mydb.DB("seat")
	service.Insert(newSTItem("testsunyetsununiversity", 1080))
	fmt.Println("Seat database init!")
}

// addSignoutTrigger 添加签退触发器，每小时的30min生效一次
func addSignoutTrigger(school string) {
	timetrigger.New(getToday(time.Now(), 0, 30), time.Hour, func() {
		// 更改这一小时的座位信息
		nowhour := getCurrentTimeInterval(time.Now())
		items := service.FindBySchoolAndTimeInterval(school, nowhour)
		nexthour := nowhour
		nexthour.AddOneHour()
		for i := 0; i < len(items); i++ {
			if items[i].Seatinfo == Book {
				user.PunishStudent(items[i].StudentID)
				items[i].Seatinfo = BookAndUnSignin
				unbookAllAfterSeat(school, nexthour, items[i].SeatID, items[i].StudentID)
			} else if items[i].Seatinfo == Signin {
				items[i].Seatinfo = Signout
				signoutAllAfterSeat(school, nexthour, items[i].SeatID, items[i].StudentID)
			}
		}
		service.UpdateAllSeat(school, nowhour, items)
		fmt.Println("Signout Trigger run:", time.Now())
	}).Run()
}

// addSeatUpdateTrigger 添加座位更新触发器，每天更新一次
func addSeatUpdateTrigger(school string) {
	timetrigger.New(getToday(time.Now(), 0, 0), 24*time.Hour, func() {
		// 添加新的座位，删除旧的座位
		service.Insert(newSTItem(school, 1080))
		deletetime := getCurrentTimeInterval(time.Now().Add(-1 * 30 * 24 * time.Hour))
		service.DeleteOldTimeInterval(school, deletetime)
		fmt.Println("Seat Update Trigger run:", time.Now())
	}).Run()
}

// addTestTrigger
func addTestTrigger() {
	timetrigger.New(getToday(time.Now(), 0, 31), time.Minute, func() {
		fmt.Println("Trigger run:")
		fmt.Println(time.Now())
	}).Run()
}
