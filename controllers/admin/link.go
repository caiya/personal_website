package admin

import (
	"personal_website/controllers"
)

type LinkHandle struct {
	controllers.BaseHandle
}

//链接首页
func (this *LinkHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/link.html", "c10")
	return
}

//新增页
func (this *LinkHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_link.html", "c10")
	return
}
