package models

import (
	"github.com/astaxie/beego/orm"
)

type Setting struct {
	Id                 int    `json:"id" form:"id"`
	Indexcount         int    `json:"indexcount" form:"indexcount" valid:"Required;Range(0,1000)"`
	Gallerycount       int    `json:"gallerycount" form:"gallerycount" valid:"Required;Range(0,1000)`
	Blogcount          int    `json:"blogcount" form:"blogcount" valid:"Required;Range(0,1000)`
	Typecount          int    `json:"typecount" form:"typecount" valid:"Required;Range(0,1000)`
	Recentpostcount    int    `json:"recentpostcount" form:"recentpostcount" valid:"Required;Range(0,1000)`
	Recentcommentcount int    `json:"recentcommentcount" form:"recentcommentcount" valid:"Required;Range(0,1000)`
	Blackip            string `json:"blackip" form:"blackip"`
	Sitename           string `json:"sitename" form:"sitename"`
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
