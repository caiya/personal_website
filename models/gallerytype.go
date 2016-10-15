package models

import (
	"github.com/astaxie/beego/orm"
)

type Gallerytype struct {
	Id       int        `json:"id" form:"id"`
	Name     string     `json:"name" form:"name"`
	Orderno  int        `json:"orderno" form:"orderno"`
	Remark   string     `json:"remark" form:"remark"`
	Gallerys []*Gallery `json:"gallerys" orm:"reverse(many)"`
	Selected int        `json:"selected" orm:"-"`
}

func (gallerytype *Gallerytype) TableName() string {
	return "gallerytype"
}

func (gallerytype *Gallerytype) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(gallerytype)
}

//注册模型
func init() {
	orm.RegisterModel(new(Gallerytype))
}

//获取所有分类
func (gallerytype *Gallerytype) GetAll() []Gallerytype {
	var gts []Gallerytype
	if _, err := gallerytype.Query().OrderBy("-Id").All(&gts); err != nil {
		return nil
	}
	return gts
}

//新增
func (gallerytype *Gallerytype) Insert() error {
	if _, err := orm.NewOrm().Insert(gallerytype); err != nil {
		return err
	}
	return nil
}

//删除
func (gallerytype *Gallerytype) Delete() error {
	if _, err := orm.NewOrm().Delete(gallerytype); err != nil {
		return err
	}
	return nil
}

//修改
func (gallerytype *Gallerytype) Update(field ...string) error {
	if _, err := orm.NewOrm().Update(gallerytype, field...); err != nil {
		return err
	}
	return nil
}

//查询
func (gallerytype *Gallerytype) Read(field ...string) error {
	if err := orm.NewOrm().Read(gallerytype, field...); err != nil {
		return err
	}
	return nil
}
