package admin

import (
	"fmt"
	"personal_website/models"
)

type IndexHandle struct {
	baseHandle
}

func (this *IndexHandle) Index() {
	this.TplName = "admin/login.html"
}

func (this *IndexHandle) Login() {
	u := Admin{Name: this.GetString("username"), Pass: this.GetString("password")}
	this.Ctx.WriteString("欢迎登录成功")
}
