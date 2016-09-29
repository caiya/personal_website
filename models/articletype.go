package models

import (
	"github.com/astaxie/beego/orm"
)

type Articletype struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Orderno  int        `json:"orderno"`
	Articles []*Article `json:"articles" orm:"reverse(many)"`
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

//新增
func (articletype *Articletype) Insert() error {
	if _, err := orm.NewOrm().Insert(articletype); err != nil {
		return err
	}
	return nil
}

//删除
func (articletype *Articletype) Delete() error {
	if _, err := orm.NewOrm().Delete(articletype); err != nil {
		return err
	}
	return nil
}

//修改
func (articletype *Articletype) Update(field ...string) error {
	if _, err := orm.NewOrm().Update(articletype, field...); err != nil {
		return err
	}
	return nil
}

//查询
func (articletype *Articletype) Read(field ...string) error {
	if err := orm.NewOrm().Read(articletype, field...); err != nil {
		return err
	}
	return nil
}

//查询所有类别
func (articletype *Articletype) GetAll() []Articletype {
	var ats []Articletype
	if _, err := articletype.Query().OrderBy("-id").All(&ats, "Id", "Name", "Orderno"); err != nil {
		return nil
	}
	return ats
}
