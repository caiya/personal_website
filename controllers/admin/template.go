package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type TemplateHandle struct {
	controllers.BaseHandle
}

func (this *TemplateHandle) Index() {
	temp := &models.Template{}
	temps, _ := temp.GetList(nil, nil)
	this.Data["datas"] = temps
	this.RspTemp("admin/layout.html", "admin/page.html", "c2")
	return
}

func (this *TemplateHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_page.html", "c2")
	return
}

//新增
func (this *TemplateHandle) Add() {
	temp := models.Template{}
	if err := this.ParseForm(&temp); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", temp)
	if temp.Id == 0 { //新增
		temp.Insert()
	} else { //修改
		temp.Update()
	}
	this.Redirect("/admin/template", 302)
}

//修改页
func (this *TemplateHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	temp := &models.Template{Id: id}
	temp.Read()
	this.Data["temp"] = temp
	this.RspTemp("admin/layout.html", "admin/add_page.html", "c2")
}

//单个删除
func (this *TemplateHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	temp := &models.Template{Id: id}
	temp.Delete()
	this.Redirect("/admin/template", 302)
}
