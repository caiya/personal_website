package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type Template struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Remark  string `json:"remark"`
}

func (template *Template) TableName() string {
	return "template"
}

//注册模型
func init() {
	orm.RegisterModel(new(Template))
}

//qs
func (template *Template) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(template)
}

//修改
func (template *Template) ToUpdate(field ...string) error {
	if _, err := orm.NewOrm().Update(template, field...); err != nil {
		return err
	}
	return nil
}

//删除
func (template *Template) Delete() error {
	if _, err := orm.NewOrm().Delete(template); err != nil {
		return err
	}
	return nil
}

//查询
func (template *Template) GetList(query map[string]string, fields []string) ([]*Template, error) {
	qs := template.Query()
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1) //-1代表全部替换
		qs = qs.Filter(k, v)
	}
	var temps []*Template
	if _, err := qs.All(&temps, fields...); err != nil {
		return nil, err
	} else {
		return temps, nil
	}
}
