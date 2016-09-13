package admin

import (
	"personal_website/controllers"
)

type TemplateHandle struct {
	controllers.BaseHandle
}

func (this *TemplateHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/page.html", "c2")
	return
}
