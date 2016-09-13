package admin

import (
	"personal_website/controllers"
)

type ManagerHandle struct {
	controllers.BaseHandle
}

//后台用户首页
func (this *ManagerHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/manager.html", "c8")
	return
}

//新增页
func (this *ManagerHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
	return
}

//修改页
func (this *ManagerHandle) Update() {

}
