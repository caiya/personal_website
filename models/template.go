package models

import (
	"github.com/astaxie/beego/orm"
)

type Template struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (template *Template) TableName() string {
	return "template"
}

//注册模型
func init() {
	orm.RegisterModel(new(Template))
}

func (template *Template) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(template)
}
