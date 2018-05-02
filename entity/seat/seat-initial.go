package seat

import (
	"fmt"

	"github.com/book-library-seat-system/go-server/orm"
	. "github.com/book-library-seat-system/go-server/util"
)

func init() {
	err := orm.Mydb.Sync2(new(Meeting))
	CheckErr(err)
	fmt.Println("Meeting database init")
}
