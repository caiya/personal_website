package admin

import (
	"personal_website/controllers"
)

type SettingHandle struct {
	controllers.BaseHandle
}

func (this *SettingHandle) Index() {
	this.Layout = "admin/index.html"
	this.TplName = "admin/system.html"
}
