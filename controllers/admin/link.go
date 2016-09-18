package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

const PAGESIZE int64 = 10

type LinkHandle struct {
	controllers.BaseHandle
}

//链接首页
func (this *LinkHandle) Index() {
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	link := &models.Link{}

	records, count, err := link.GetList(nil, nil, []string{"id"}, []string{"desc"}, int64(page), PAGESIZE)
	if err == nil {
		this.Data["datas"] = records //记录数
		this.Data["count"] = count   //总记录数
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
	this.Data["pageCount"] = pageCount //总页数
	this.Data["page"] = page           //当前页
	this.RspTemp("admin/layout.html", "admin/link.html", "c10")
	return
}

//新增页
func (this *LinkHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_link.html", "c10")
	return
}

//单个删除
func (this *LinkHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	link := &models.Link{Id: id}
	link.Delete()
	this.Redirect("/admin/link", 302)
}

//批量删除
func (this *LinkHandle) Del() {
	link := &models.Link{}
	idstr := this.GetString("ids")
	if err := link.Del(idstr); err != nil {
		this.RspJson(false, err.Error())
		return
	}
	this.RspJson(true, "删除成功")
}

//新增链接
func (this *LinkHandle) Add() {
	fmt.Println("请求过来了.............")
	link := models.Link{}
	if err := this.ParseForm(&link); err != nil {
		fmt.Println(err.Error())
		return
	}
	link.Insert()
	this.Redirect("/admin/link", 302)
}
