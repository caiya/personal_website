package admin

import (
	"personal_website/controllers"
)

type GallerytypeHandle struct {
	controllers.BaseHandle
}

//照片分类首页
func (this *GallerytypeHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/gallery_category.html", "c3")
	return
}

//新增页
func (this *GallerytypeHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_gallery_category.html", "c3")
	return
}
