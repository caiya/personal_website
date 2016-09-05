package models

import (
	"github.com/astaxie/beego/orm"
)

type Link struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Orderno int    `json:"orderno"`
}

func (link *Link) TableName() string {
	return "link"
}

func (link *Link) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(link)
}
