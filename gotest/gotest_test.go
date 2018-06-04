package gotest

import (
	//"time"
	"os"
	"testing"

	"github.com/book-library-seat-system/go-server/entity/seat"
	"github.com/book-library-seat-system/go-server/entity/user"
)

// 测试前
func TestBegin(t *testing.T) {
	// ser := server.NewServer()
	// ser.Run(":8899")
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
	seat.TestDeleteOldTimeInterval(t)
	seat.TestDeleteBySchool(t)
}

// UserHandle 测试部分
// func TestUserHandle(t *testing.T) {
// 	server.TesttestGET(t)
// 	server.TesttestPost(t)
// }

// 测试完毕
func TestEnd(t *testing.T) {
	os.Exit(0)
}
