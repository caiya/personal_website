package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Nickname  string     `json:"nickname"`
	Mobile    string     `json:"mobile"`
	Age       int        `json:"age"`
	Sex       bool       `json:"sex"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	Pass      string     `json:"pass"`
	Addtime   int        `json:"addtime"`
	Lastlogin int        `json:"lastlogin"`
	Articles  []*Article `orm:"reverse(many)"`
	Isadmin   int
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(user)
}

//注册模型
func init() {
	orm.RegisterModel(new(User))
}
