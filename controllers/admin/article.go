package admin

import (
	"personal_website/controllers"
)

type ArticleHandle struct {
	controllers.BaseHandle
}

//博客首页
func (this *ArticleHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/article.html", "c6")
	return
}

//新增页
func (this *ArticleHandle) Add() {
	this.RspTemp("admin/layout.html", "admin/add_article.html", "c6")
	return
}
