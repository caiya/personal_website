package admin

import (
	"personal_website/controllers"
)

type SettingHandle struct {
	controllers.BaseHandle
}

//设置页
func (this *SettingHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/system.html", "c1")
	return
}
