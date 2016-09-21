package models

import (
	"github.com/astaxie/beego/orm"
)

type Articletype struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Orderno  int        `json:"orderno"`
	Articles []*Article `orm:"reverse(many)"`
}

func (articletype *Articletype) TableName() string {
	return "articletype"
}

func (articletype *Articletype) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(articletype)
}

//注册模型
func init() {
	orm.RegisterModel(new(Articletype))
}
