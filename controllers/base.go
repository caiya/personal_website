package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	fmt.Println(controller)

	//读取配置文件
	if len(CONFIG) == 0 {
		cfg, _ := ini.Load("./conf/config.ini")
		val := cfg.Section("whitecontroller").Key("controller").String()
		CONFIG = strings.Split(val, ",")
	}
	var flag = false
	if u == nil { //如果用户未登录并且访问非白名单
		for _, v := range CONFIG {
			if v == controller || strings.Contains(controller, "Controller") { //在白名单，直接退出循环
				flag = true
				break
			}
		}
		if !flag {
			this.Redirect("/admin", 302)
			return
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

//前台模板页输出
func (this *BaseHandle) RenderHtml(layout string, tplname string, objstr []string) {
	for _, value := range objstr {
		v := strings.Split(value, "_")
		this.Data[v[0]] = v[1]
	}
	fmt.Println(objstr)
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

//上传图片
func (this *BaseHandle) UploadImg() {
	f, _, err := this.GetFile("editormd-image-file")
	defer f.Close()
	if err != nil {
		fmt.Println("get file error ", err.Error())
		this.Data["json"] = &map[string]interface{}{"url": "", "success": 0, "msg": err.Error()}
	} else {
		imgurl := strconv.Itoa(int(time.Now().Unix())) + ".png"
		this.SaveToFile("editormd-image-file", "static/upload/"+imgurl)
		this.Data["json"] = &map[string]interface{}{"url": "/static/upload/" + imgurl, "success": 1, "msg": "图片上传成功"}
	}
	this.ServeJSON()
}
