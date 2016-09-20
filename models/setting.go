package models

import (
	"github.com/astaxie/beego/orm"
)

type Setting struct {
	Id                 int    `json:"id"`
	Indexcount         int    `json:"indexcount"`
	Gallerycount       int    `json:"gallerycount"`
	Blogcount          int    `json:"blogcount"`
	Typecount          int    `json:"typecount"`
	Recentpostcount    int    `json:"recentpostcount"`
	Recentcommentcount int    `json:"recentcommentcount"`
	Blackip            string `json:"blackip"`
	Sitename           string `json:"sitename"`
}

func (setting *Setting) TableName() string {
	return "setting"
}

func (setting *Setting) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(setting)
}

//注册数据库模型
func init() {
	orm.RegisterModel(new(Setting))
}

//新增
func (setting *Setting) Insert() error {
	if _, err := orm.NewOrm().Insert(setting); err != nil {
		return err
	}
	return nil
}

//删除
func (setting *Setting) Delete() error {
	if _, err := orm.NewOrm().Delete(setting); err != nil {
		return err
	}
	return nil
}

//修改
func (setting *Setting) Update(field ...string) error {
	if _, err := orm.NewOrm().Update(setting, field...); err != nil {
		return err
	}
	return nil
}

//查询
func (setting *Setting) Read(field ...string) error {
	if err := orm.NewOrm().Read(setting, field...); err != nil {
		return err
	}
	return nil
}
