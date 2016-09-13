package admin

import (
	"personal_website/controllers"
)

type GalleryHandle struct {
	controllers.BaseHandle
}

//照片首页
func (this *GalleryHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/gallery.html", "c4")
	return
}

//新增页
func (this *GalleryHandle) Add() {
	this.RspTemp("admin/layout.html", "admin/add_gallery.html", "c4")
	return
}
