package routers

import (
	"personal_website/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {

	//用户登录退出
	beego.Router("/admin", &admin.IndexHandle{}, "GET:Index")
	beego.Router("/admin/login", &admin.IndexHandle{}, "POST:Login")
	beego.Router("/admin/exit", &admin.IndexHandle{}, "GET:Exit")
	beego.Router("/admin/main", &admin.MainHandle{}, "GET:Main")

	//跳转路由
	beego.Router("/admin/setting", &admin.SettingHandle{}, "GET:Index")

	beego.Router("/admin/template", &admin.TemplateHandle{}, "GET:Index")
	beego.Router("/admin/template/add", &admin.TemplateHandle{}, "GET:Add")

	beego.Router("/admin/gallerytype", &admin.GallerytypeHandle{}, "GET:Index")
	beego.Router("/admin/gallerytype/add", &admin.GallerytypeHandle{}, "GET:Add")

	beego.Router("/admin/gallery", &admin.GalleryHandle{}, "GET:Index")
	beego.Router("/admin/gallery/add", &admin.GalleryHandle{}, "GET:Add")

}
