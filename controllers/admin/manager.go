package admin

import (
	"personal_website/controllers"
)

type UseradminHandle struct {
	controllers.BaseHandle
}

//后台用户首页
func (this *UseradminHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/manager.html", "c8")
	return
}

//新增页
func (this *UseradminHandle) Add() {
	this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
	return
}
