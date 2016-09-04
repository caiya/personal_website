package models

import (
	"time"
)

type Article struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Addtime  time.Time `json:"addtime"`
	Uptime   time.Time `json:"uptime"`
	Uid      int       `json:"uid"`
	Uname    string    `json:"uname"` //作者名称-非数据库字段
	Link     string    `json:"link"`
	Intro    string    `json:"intro"`
	Typeid   int       `json:"typeid"`
	Typename string    `json:"typename"`
}

func (article *Article) TableName() string {
	return "article"
}
