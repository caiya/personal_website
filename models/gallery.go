package models

import (
	"github.com/astaxie/beego/orm"
)

type Gallery struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Imgurl  string `json:"imgurl"`
	Orderno int    `json:"orderno"`
}

func (gallery *Gallery) TableName() string {
	return "gallery"
}

func (gallery *Gallery) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(gallery)
}
