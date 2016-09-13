package controllers

import (
	"github.com/astaxie/beego"
)

type BaseHandle struct {
	beego.Controller
}

//json输出
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
	this.Data["currUser"] = u
}

//后台模板页输出
func (this *BaseHandle) RspTemp(layout string, tplname string, class string) {
	if class != "" {
		this.Data[class] = "cur"
	}
	this.Layout = layout
	this.TplName = tplname
}

//获取客户机ip
func (this *BaseHandle) GetIP() string {
	ip := this.Ctx.Request.Header.Get("X-Forwarded-For") //反向代理获取ip
	//this.Ctx.Request.RemoteAddr 直接获取ip
	if ip != "" {
		return ip
	} else {
		return this.Ctx.Request.RemoteAddr
	}
}
