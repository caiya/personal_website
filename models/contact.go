package models

import (
	"github.com/astaxie/beego/orm"
)

type Contact struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Message string `json:"message"`
	Addtime int    `json:"addtime"`
}

func (contact *Contact) TableName() string {
	return "contact"
}

func (contact *Contact) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(contact)
}
