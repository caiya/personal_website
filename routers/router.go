package routers

import (
	"personal_website/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	//后台 begin

	//用户登录退出
	beego.Router("/admin", &admin.IndexHandle{}, "GET:Index")
	beego.Router("/admin/login", &admin.IndexHandle{}, "POST:Login")
	beego.Router("/admin/exit", &admin.IndexHandle{}, "GET:Exit")
	beego.Router("/admin/main", &admin.MainHandle{}, "*:Main")

	//系统设置
	beego.Router("/setting", &admin.SettingHandle{}, "*:Index")
}
