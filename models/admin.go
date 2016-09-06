package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
}

func (admin *Admin) TableName() string {
	return "admin"
}

func (admin *Admin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(admin)
}

//新增
func (admin *Admin) Insert() error {
	if _, err := orm.NewOrm().Insert(admin); err != nil {
		return err
	}
	return nil
}

//删除
func (admin *Admin) Delete() error {
	if _, err := orm.NewOrm().Delete(admin); err != nil {
		return err
	}
	return nil
}

//修改
func (admin *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(admin, fields); err != nil {
		return err
	}
	return nil
}

//查询
func (admin *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(admin); err != nil {
		return err
	}
	return nil
}

//分页查询
func (admin *Admin) GetAdminList(query map[string]string, fields []string, sortby []string, orderby []string, offcet int, limit int) (ml []interface{}, err error) {
	var ad Admin
	qs := ad.Query()
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
}
