package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
	"time"
)

type ArticleHandle struct {
	controllers.BaseHandle
}

//博客首页
func (this *ArticleHandle) Index() {
	id := this.GetString("id")
	title := this.GetString("title")

	query := make(map[string]string, 0)

	if id != "" {
		query["id"] = id
	}
	if title != "" {
		query["title.contains"] = title
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
	fmt.Println(articles)
	//数据回显
	this.Data["id"] = id
	this.Data["title"] = title
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
	article := &models.Article{}
	if err := this.ParseForm(article); err != nil {
		fmt.Println(err.Error())
		return
	}
	t, _ := this.GetInt("articletype")
	article.Articletype = &models.Articletype{Id: t}
	article.User = &models.User{Id: 1} //指定当前用户为admin用户
	now := int(time.Now().Unix())
	article.Uptime = now
	if article.Id == 0 { //新增
		article.Addtime = now
		article.User = &models.User{Id: 1} //指定当前用户为admin用户
		article.Insert()
	} else { //修改
		article.Update("Id", "Title", "Content", "Uptime", "User", "Link", "Intro", "Articletype")
	}
	this.Redirect("/admin/article", 302)
}

//修改页
func (this *ArticleHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article := &models.Article{Id: id}
	article.Read()

	articletype := &models.Articletype{}
	articletypes := articletype.GetAll()
	this.Data["articletypes"] = articletypes

	fmt.Println(article.Articletype.Id)

	for index, value := range articletypes {
		if value.Id == article.Articletype.Id {
			articletypes[index].Selected = 1
		} else {
			articletypes[index].Selected = 0
		}
	}
	this.Data["article"] = article
	this.RspTemp("admin/layout.html", "admin/add_article.html", "c6")
}

//删除
func (this *ArticleHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article := &models.Article{Id: id}
	article.Delete()
	this.Redirect("/admin/article", 302)
}
