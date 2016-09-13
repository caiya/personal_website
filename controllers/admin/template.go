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

func (this *TemplateHandle) Add() {
	this.RspTemp("admin/layout.html", "admin/addpage.html", "c2")
	return
}
