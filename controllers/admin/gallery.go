package admin

import (
	"personal_website/controllers"
	"strconv"
)

type GalleryHandle struct {
	controllers.BaseHandle
}

//照片首页
func (this *GalleryHandle) Index() {
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	this.Data["page"] = page
	this.Data["pageCount"] = 1 //总页数
	this.RspTemp("admin/layout.html", "admin/gallery.html", "c4")
}

//新增页
func (this *GalleryHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_gallery.html", "c4")
}
