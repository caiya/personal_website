package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type ArticletypeHandle struct {
	controllers.BaseHandle
}

//博客分类首页
func (this *ArticletypeHandle) Index() {
	articletype := &models.Articletype{}
	articletypes := articletype.GetAll()
	this.Data["datas"] = articletypes
	this.RspTemp("admin/layout.html", "admin/article_category.html", "c5")
	return
}

//新增页
func (this *ArticletypeHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_article_category.html", "c5")
}

//修改页
func (this *ArticletypeHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	articletype := &models.Articletype{Id: id}
	articletype.Read()
	this.Data["articletype"] = articletype
	this.RspTemp("admin/layout.html", "admin/add_article_category.html", "c5")
}

//新增|修改操作
func (this *ArticletypeHandle) Add() {
	articletype := models.Articletype{}
	if err := this.ParseForm(&articletype); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", articletype)
	if articletype.Id == 0 { //新增
		articletype.Insert()
	} else { //修改
		articletype.Update()
	}
	this.Redirect("/admin/articletype", 302)
}

//删除
func (this *ArticletypeHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	articletype := &models.Articletype{Id: id}
	articletype.Delete()
	this.Redirect("/admin/articletype", 302)
}
