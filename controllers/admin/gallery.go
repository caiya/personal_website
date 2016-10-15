package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

const GALLERYSIZE = 8

type GalleryHandle struct {
	controllers.BaseHandle
}

//照片首页
func (this *GalleryHandle) Index() {
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	gallery := new(models.Gallery)

	gallerys, count, err := gallery.GetList(nil, nil, []string{"id"}, []string{"desc"}, int64(page), GALLERYSIZE)
	if err == nil {
		this.Data["datas"] = gallerys //记录数
		this.Data["count"] = count    //总记录数
	} else {
		this.Data["datas"] = nil
		this.Data["count"] = 0
	}
	var pageCount int64
	if count%GALLERYSIZE == 0 {
		pageCount = count / GALLERYSIZE
	} else {
		pageCount = count/GALLERYSIZE + 1
	}
	fmt.Println(gallerys)
	this.Data["pageCount"] = pageCount //总页数
	this.Data["page"] = page           //当前页
	this.RspTemp("admin/layout.html", "admin/gallery.html", "c4")
}

//新增页
func (this *GalleryHandle) ToAdd() {
	gallerytype := new(models.Gallerytype)
	gallerytypes := gallerytype.GetAll()
	this.Data["gallerytypes"] = gallerytypes
	this.RspTemp("admin/layout.html", "admin/add_gallery.html", "c4")
}

//修改页
func (this *GalleryHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	gallery := &models.Gallery{Id: id}
	gallery.Read()
	this.Data["gallery"] = gallery

	gallerytype := new(models.Gallerytype)
	gallerytypes := gallerytype.GetAll()

	fmt.Println(gallerytypes)

	for index, value := range gallerytypes {
		if value.Id == gallery.Gallerytype.Id {
			gallerytypes[index].Selected = 1
		} else {
			gallerytypes[index].Selected = 0
		}
	}

	this.Data["gallerytypes"] = gallerytypes
	this.RspTemp("admin/layout.html", "admin/add_gallery.html", "c4")
}

//新增
func (this *GalleryHandle) Add() {

	gallery := new(models.Gallery)
	if err := this.ParseForm(gallery); err != nil {
		fmt.Println(err.Error())
		return
	}
	t, _ := this.GetInt("gallerytype")
	gallery.Gallerytype = &models.Gallerytype{Id: t}
	if gallery.Id == 0 { //新增
		gallery.Insert()
	} else { //修改
		gallery.Update()
	}
	this.Redirect("/admin/gallery", 302)
}
