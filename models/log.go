package models

import (
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id     int    `json:"id"`
	Optime int    `json:"optime"`
	Opid   int    `json:"opid"`
	Opname string `json:"opname"`
	Record string `json:"record"`
}

func (record *Record) TableName() string {
	return "log"
}

func (record *Record) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(record)
}

//注册模型
func init() {
	orm.RegisterModel(new(Record))
}
