package gotest

import (
	//"time"
	"testing"

	"github.com/book-library-seat-system/go-server/entity/seat"
	"github.com/book-library-seat-system/go-server/entity/user"
)

// 测试前
func TestBegin(t *testing.T) {
}

// user测试部分
func TestUserDao(t *testing.T) {
	user.TestSave(t)
	user.TestFindByID(t)
	user.TestFindSchoolByID(t)
	user.TestUpdate(t)
	user.TestDeleteByID(t)
}

// seat测试部分
func TestSeatDao(t *testing.T) {
	seat.TestInsert(t)
	// seat.TestFindAll(t)
	seat.TestFindBySchool(t)
	seat.TestFindBySchoolAndTimeInterval(t)
	seat.TestFindOneSeat(t)
	seat.TestUpdateAllSeat(t)
	seat.TestUpdateOneSeat(t)
	seat.TestDeleteBySchoolAndTimeInterval(t)
	seat.TestDeleteBySchool(t)
}

/*
func TestCreateStudent(t *testing.T){
	var userTest user.Item
	userTest.ID = "111"
	userTest.NetID = "hmy"
	userTest.Hashpassword = "123"
	userTest.School = "sysu"
	userTest.Violation = 0
	err := user.RegisterStudent(userTest)
	if err != nil {
		t.Errorf("error:%s", err)
	}  
}

func TestListStudent(t *testing.T){

	var openID = "111"
	err := user.GetStudent(openID)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}

func TestShowTimeInterval(t *testing.T){
	var openID = "111"
    school = user.GetStudentsSchool(openID)
	err := seat.GetAllTimeInterval(school)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}


var t1 = time.Now()
var t2 = t1.Add(time.Hour).Add(time.Hour)
var school = "sysu"
func TestShowSeat(t *testing.T){
	var timeInterval seat.TimeInterval
	timeInterval.Begintime = t1
	timeInterval.Endtime = t2
	err := seat.GetAllSeatinfo(school, timeInterval)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}

func TestBookSeat(t *testing.T){
	var timeInterval seat.TimeInterval
	timeInterval.Begintime = t1
	timeInterval.Endtime = t2
	studentid = "15331116"
	seatid = 1
	err := seat.BookSeat(school, timeInterval, studentid, seatid)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}

func TestUnBookSeat(t *testing.T){
	var timeInterval seat.TimeInterval
	timeInterval.Begintime = t1
	timeInterval.Endtime = t2
	var studentid = "15331116"
	var seatid = 1
	err := seat.UnbookSeat(school, timeInterval, studentid, seatid)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}

func TestSigninSeat(t *testing.T){
	var studentid = "15331116"
	var seatid = 1
	err := seat.SigninSeat(school, studentid, seatid)
	if err != nil {
		t.Errorf("error:%s", err)
	} 
}
*/
// 测试完毕
func TestEnd(t *testing.T) {
	
}
