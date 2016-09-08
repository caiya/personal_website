package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
)

type IndexHandle struct {
	controllers.BaseHandle
}

func (this *IndexHandle) Index() {
	this.TplName = "admin/login.html"
}

func (this *IndexHandle) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	u := &models.Admin{Name: username}
	u.Read("Name")
	if u.Id == 0 {

	}
	this.Ctx.WriteString("欢迎登录成功")
}
