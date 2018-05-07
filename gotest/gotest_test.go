package gotest

import (
	"testing"

	"github.com/book-library-seat-system/go-server/entity/user"
)

// 测试前，清空数据库
func TestBegin(t *testing.T) {

}

// user测试部分
func TestUserDao(t *testing.T) {
	user.TestSave(t)
	user.TestFindByID(t)
	user.TestUpdate(t)
	user.TestDeleteByID(t)
	user.TestShow(t)
}

// var t1 = time.Now()
// var t2 = t1.Add(time.Second).Add(time.Second)
// var startTimeConflict = t1.Add(time.Second)
// var endTimeConflict = t2.Add(time.Second)
// var meetName = t1.Format("2006-01-02 15:04:05")
// var title = "three persons' team"

// 测试完毕，清空数据库
func TestEnd(t *testing.T) {
}
