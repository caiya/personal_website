package routers

import (
	"personal_website/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/*", &admin.IndexHandle{}, "GET:Index")
	beego.Router("/admin/login", &admin.IndexHandle{}, "POST:Login")
	beego.Router("/admin/main", &admin.IndexHandle{}, "*:Main")
}
