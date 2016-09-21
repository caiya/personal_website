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
	Uid      *User        `json:"uid" orm:"rel(fk)"`
	Uname    string       `json:"uname"` //作者名称-非数据库字段
	Link     string       `json:"link"`
	Intro    string       `json:"intro"`
	Typeid   *Articletype `json:"typeid" orm:"rel(fk)"`
	Typename string       `json:"typename"`
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
