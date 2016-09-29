package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type ArticleHandle struct {
	controllers.BaseHandle
}

//博客首页
func (this *ArticleHandle) Index() {
	id := this.GetString("id")
	name := this.GetString("name")
	url := this.GetString("url")

	query := make(map[string]string, 0)

	if id != "" {
		query["id"] = id
	}
	if name != "" {
		query["name.contains"] = name
	}
	if url != "" {
		query["url.contains"] = url
	}

	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	article := &models.Article{}

	articles, count, err := article.GetList(query, nil, []string{"id"}, []string{"desc"}, int64(page), PAGESIZE)
	if err == nil {
		this.Data["datas"] = articles //记录数
		this.Data["count"] = count    //总记录数
	} else {
		this.Data["datas"] = nil
		this.Data["count"] = 0
	}
	var pageCount int64
	if count%PAGESIZE == 0 {
		pageCount = count / PAGESIZE
	} else {
		pageCount = count/PAGESIZE + 1
	}
	//数据回显
	this.Data["id"] = id
	this.Data["name"] = name
	this.Data["url"] = url
	this.Data["pageCount"] = pageCount //总页数
	this.Data["page"] = page           //当前页
	this.RspTemp("admin/layout.html", "admin/article.html", "c6")
	return
}

//新增页
func (this *ArticleHandle) ToAdd() {
	//获取文章分类下拉框数据
	articletype := &models.Articletype{}
	articletypes := articletype.GetAll()
	this.Data["articletypes"] = articletypes
	this.RspTemp("admin/layout.html", "admin/add_article.html", "c6")
	return
}

//新增
func (this *ArticleHandle) Add() {

}

//修改页
func (this *ArticleHandle) ToUpdate() {

}

//修改
func (this *ArticleHandle) Update() {

}
