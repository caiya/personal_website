package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id      int    `json:"id"`
	Cid     int    `json:"cid"`
	Cname   string `json:"cname"`
	Cemail  string `json:"cemail"`
	Content string `json:"content"`
	Addtime int    `json:"addtime"`
	Aid     int    `json:"aid"`
}

func (comment *Comment) TableName() string {
	return "comment"
}

func (comment *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(comment)
}
