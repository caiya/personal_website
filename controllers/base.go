package controllers

import (
	"github.com/astaxie/beego"
)

type BaseHandle struct {
	beego.Controller
}

func (this *BaseHandle) RspJson(status bool, msg string) {
	this.Data["json"] = &map[string]interface{}{"succ": status, "msg": msg}
	this.ServeJSON()
}

//登录校验
func (this *BaseHandle) Prepare() {
	u := this.GetSession("currUser")
	controller, action := this.GetControllerAndAction()
	if u == nil && controller != "IndexHandle" && (action != "Index" || action != "Login") {
		this.Redirect("/admin", 302)
		return
	}
}
