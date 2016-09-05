package models

import (
	"github.com/astaxie/beego/orm"
)

type Articletype struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Orderno int    `json:"orderno"`
}

func (articletype *Articletype) TableName() string {
	return "articletype"
}

func (articletype *Articletype) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(articletype)
}
