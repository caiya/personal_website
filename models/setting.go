package models

import (
	"github.com/astaxie/beego/orm"
)

type Setting struct {
	Id                 int `json:"id"`
	Indexcount         int `json:"indexcount"`
	Gallerycount       int `json:"gallerycount"`
	Blogcount          int `json:"blogcount"`
	Typecount          int `json:"typecount"`
	Recentpostcount    int `json:"recentpostcount"`
	Recentcommentcount int `json:"recentcommentcount"`
}

func (setting *Setting) TableName() string {
	return "setting"
}

func (setting *Setting) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(setting)
}
