package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type GallerytypeHandle struct {
	controllers.BaseHandle
}

//照片分类首页
func (this *GallerytypeHandle) Index() {
	gallerytype := &models.Gallerytype{}
	gallerytypes := gallerytype.GetAll()
	this.Data["datas"] = gallerytypes

	fmt.Println(gallerytypes)

	this.RspTemp("admin/layout.html", "admin/gallery_category.html", "c3")
	return
}

//新增页
func (this *GallerytypeHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_gallery_category.html", "c3")
	return
}

//修改页
func (this *GallerytypeHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	gallerytype := &models.Gallerytype{Id: id}
	gallerytype.Read()
	this.Data["gallerytype"] = gallerytype
	this.RspTemp("admin/layout.html", "admin/add_gallery_category.html", "c3")
}

//新增|修改操作
func (this *GallerytypeHandle) Add() {
	gallerytype := models.Gallerytype{}
	if err := this.ParseForm(&gallerytype); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", gallerytype)
	if gallerytype.Id == 0 { //新增
		gallerytype.Insert()
	} else { //修改
		gallerytype.Update()
	}
	this.Redirect("/admin/gallerytype", 302)
}

//删除
func (this *GallerytypeHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	gallerytype := &models.Gallerytype{Id: id}
	gallerytype.Delete()
	this.Redirect("/admin/gallerytype", 302)
}
