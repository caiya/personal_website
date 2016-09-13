package models

import (
	"github.com/astaxie/beego/orm"
)

type Gallerytype struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Orderno int    `json:"orderno"`
	Desc    string `json:"desc"`
}

func (gallerytype *Gallerytype) TableName() string {
	return "gallerytype"
}

func (gallerytype *Gallerytype) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(gallerytype)
}
