package orm

import (
	"fmt"
	"path/filepath"

	. "github.com/book-library-seat-system/go-server/util"
	"github.com/go-xorm/xorm"
	// 使用sqlite3数据库
	_ "github.com/mattn/go-sqlite3"
)

// Mydb 数据库指针
var Mydb *xorm.Engine

// 生成数据库，对数据库进行链接
func init() {
	// 链接sqlite3数据库
	//db, err := xorm.NewEngine("sqlite3", "./agenda-cs.db")
	databasePath := "src/github.com/book-library-seat-system/go-server/agenda-cs.db"
	databasePath = filepath.Join(*GetGOPATH(), databasePath)
	fmt.Println("Your database:", databasePath)
	db, err := xorm.NewEngine("sqlite3", databasePath)
	CheckErr(err)

	Mydb = db
}
