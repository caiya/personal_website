package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
)

type TemplateHandle struct {
	controllers.BaseHandle
}

func (this *TemplateHandle) Index() {
	temp := &models.Template{}
	temps, _ := temp.GetList(nil, nil)
	this.Data["datas"] = temps
	this.RspTemp("admin/layout.html", "admin/page.html", "c2")
	return
}

func (this *TemplateHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_page.html", "c2")
	return
}
