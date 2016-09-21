package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id      int      `json:"id"`
	Cname   string   `json:"cname"`
	Cemail  string   `json:"cemail"`
	Content string   `json:"content"`
	Addtime int      `json:"addtime"`
	Aid     *Article `json:"aid" orm:"rel(fk)"`
}

func (comment *Comment) TableName() string {
	return "comment"
}

func (comment *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(comment)
}

//注册模型
func init() {
	orm.RegisterModel(new(Comment))
}
