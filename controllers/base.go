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
