package admin

import (
	"personal_website/controllers"
	"personal_website/models"
)

type SettingHandle struct {
	controllers.BaseHandle
}

//设置页
func (this *SettingHandle) Index() {
	setting := &models.Setting{Id: 1} //读取第一条数据
	if err := setting.Read(); err != nil {
		return
	}
	if setting.Id == 0 { //未查询到数据：add -> read
		if err := setting.Insert(); err != nil {
			return
		}
		setting.Read()
	}
	this.Data["setting"] = setting
	this.RspTemp("admin/layout.html", "admin/system.html", "c1")
	return
}
