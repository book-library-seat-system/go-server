/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 工厂包，用于生产对象
Date: 2018年7月8日 星期日 下午1:50
****************************************************************************/

package seat

// newItems 生成一个Item数组，id从0开始
func newItems(seatnumber int) []Item {
	items := make([]Item, seatnumber)
	for i := 0; i < seatnumber; i++ {
		items[i].SeatID = i
		items[i].Seatinfo = UnBook
		items[i].StudentID = ""
	}
	return items
}

// newSeatInfo 生成一个新的SeatInfo
func newSeatInfo(timeinterval TimeInterval, item Item) *SeatInfo {
	seatinfo := &SeatInfo{}
	seatinfo.TimeInterval = timeinterval
	seatinfo.SeatID = item.SeatID
	seatinfo.Seatinfo = item.Seatinfo
	return seatinfo
}

// newTItems 生成一个TItem数组，timeinterval从当前时间段开始，数组数量从配置文件读取
func newTItems(seatnumber int) []TItem {
	titems := []TItem{}
	for _, timeinterval := range currentTimeIntervals() {
		titems = append(titems, TItem{
			Timeinterval: timeinterval,
			Items:        newItems(seatnumber),
		})
	}
	return titems
}

// newSTItem 生成一个STItem
func newSTItem(school string, seatnumber int) *STItem {
	newtitems := new(STItem)
	newtitems.School = school
	newtitems.Titems = newTItems(seatnumber)
	return newtitems
}
