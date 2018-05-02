package user

// Item 用户信息
type Item struct {
	// 用户ID，是唯一主键
	ID string `xorm:"pk" json:",omitempty"`
	// 用户姓名
	Name string `json"name,omitempty"`
	// hash过的密码
	Hashpassword string `json:"-"`
	// 注册用的邮箱
	Email string `json:"email,omitempty"`
	// 学生所在学校
	School string `json:"school,omitempty"`
	// 学生被警告次数
	Violation int `json:"violation"`
}

// newItem 新建一个UserItem，并返回指针
func newItem(ID string, name string, password string,
	email string, school string) *Item {
	newUserItem := new(Item)
	newUserItem.ID = ID
	newUserItem.Name = name
	newUserItem.Hashpassword = password
	newUserItem.Email = email
	newUserItem.School = school
	newUserItem.Violation = 0
	return newUserItem
}
