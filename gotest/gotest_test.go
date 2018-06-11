package gotest

import (
	//"time"

	"os/exec"
	"testing"

	"github.com/book-library-seat-system/go-server/entity/seat"
	"github.com/book-library-seat-system/go-server/entity/user"
	"github.com/book-library-seat-system/go-server/server"
)

var cmd *exec.Cmd

// TestBegin 测试前
func TestBegin(t *testing.T) {
}

// TestUserDao userdao测试部分
func TestUserDao(t *testing.T) {
	user.TestSave(t)
	user.TestFindByID(t)
	user.TestFindSchoolByID(t)
	user.TestUpdate(t)
	user.TestDeleteByID(t)
}

// TestSeatDao seatdao测试部分
func TestSeatDao(t *testing.T) {
	seat.TestInsert1(t)
	seat.TestInsert2(t)
	// seat.TestFindAll(t)
	seat.TestFindBySchool(t)
	seat.TestFindBySchoolAndTimeInterval(t)
	seat.TestFindBySchoolAndStudentID(t)
	seat.TestFindOneSeat(t)
	seat.TestUpdateAllSeat(t)
	seat.TestUpdateOneSeat(t)
	seat.TestDeleteBySchoolAndTimeInterval(t)
	seat.TestDeleteOldTimeInterval(t)
	seat.TestDeleteBySchool(t)
}

// TestUserHandle userhandle测试部分
func TestUserHandle(t *testing.T) {
	server.TesttestGET(t)
	server.TesttestPost(t)
}

// TestSeatHandle seathandle测试部分
func TestSeatHandle(t *testing.T) {

}

// TestEnd 测试完毕
func TestEnd(t *testing.T) {
}
