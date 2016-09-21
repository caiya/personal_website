package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int          `json:"id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Addtime  int          `json:"addtime"`
	Uptime   int          `json:"uptime"`
	User     *User        `json:"user" orm:"rel(fk)"`
	Link     string       `json:"link"`
	Intro    string       `json:"intro"`
	Type     *Articletype `json:"type" orm:"rel(fk)"`
	Comments []*Comment   `orm:"reverse(many)"` //反向一对多关联
}

func (article *Article) TableName() string {
	return "article"
}

func (article *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(article)
}

func init() {
	orm.RegisterModel(new(Article))
}

func (article *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(article); err != nil {
		return err
	}
	return nil
}

func (article *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(article); err != nil {
		return err
	}
	return nil
}

func (article *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(article, fields...); err != nil {
		return err
	}
	return nil
}

func (article *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(article, fields...); err != nil {
		return err
	}
	return nil
}
