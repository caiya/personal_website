package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"

	"github.com/astaxie/beego/validation"
)

type SettingHandle struct {
	controllers.BaseHandle
}

//设置页
func (this *SettingHandle) Index() {
	setting := &models.Setting{Id: 1} //读取第一条数据
	if err := setting.Read(); err != nil {
		//读取不到数据
		if err := setting.Insert(); err != nil {
			fmt.Println(err)
			return
		}
		setting.Read()
	}
	this.Data["setting"] = setting
	this.RspTemp("admin/layout.html", "admin/system.html", "c1")
	return
}

//添加|修改
func (this *SettingHandle) Update() {
	setting := &models.Setting{}
	if err := this.ParseForm(setting); err != nil {
		setting.Read("id")
		this.Data["setting"] = setting
		this.Data["errmsg"] = err.Error()
		this.RspTemp("admin/layout.html", "admin/system.html", "c1")
		return
	}
	//开始校验数据
	valid := validation.Validation{}
	b, err := valid.Valid(setting)
	if err != nil { //程序出错
		this.Data["errmsg"] = err.Error()
		setting.Read("id")
		this.Data["setting"] = setting
		this.RspTemp("admin/layout.html", "admin/system.html", "c1")
		return
	}
	if !b { //校验失败
		for _, err := range valid.Errors {
			setting.Read("id")
			this.Data["setting"] = setting
			this.Data["errmsg"] = err.Error()
			this.RspTemp("admin/layout.html", "admin/system.html", "c1")
			return
		}
	}

	if setting.Id == 0 { //新增
		setting.Insert()
	} else { //修改
		setting.Update()
	}
	setting.Read("id")
	this.Data["setting"] = setting
	this.Data["errmsg"] = "修改成功"
	this.RspTemp("admin/layout.html", "admin/system.html", "c1")
	return
}
