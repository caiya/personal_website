package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/go-ini/ini"
)

var CONFIG = make([]string, 0) //全局的配置对象

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
	controller, _ := this.GetControllerAndAction()

	//读取配置文件
	if len(CONFIG) == 0 {
		cfg, _ := ini.Load("./conf/config.ini")
		val := cfg.Section("whitecontroller").Key("controller").String()
		CONFIG = strings.Split(val, ",")
	}

	if u == nil { //如果用户未登录并且访问非白名单
		for _, val := range CONFIG {
			if val != controller {
				this.Redirect("/admin", 302)
				return
			}
		}
	} else {
		this.Data["currUser"] = u
	}
}

//后台模板页输出
func (this *BaseHandle) RspTemp(layout string, tplname string, class string) {
	if class != "" {
		this.Data[class] = "cur"
	}
	this.Layout = layout
	this.TplName = tplname
}

//获取client ip
func (this *BaseHandle) GetIP() string {
	//	ip := this.Ctx.Request.Header.Get("X-Forwarded-For") //反向代理获取ip
	//	if ip != "" {
	//		return ip
	//	} else {
	//		return this.Ctx.Request.RemoteAddr / this.Ctx.Input.IP()
	//	}
	return this.Ctx.Input.IP()
}
