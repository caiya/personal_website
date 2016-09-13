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
	beego.Router("/admin/template/add", &admin.TemplateHandle{}, "GET:ToAdd")

	beego.Router("/admin/gallerytype", &admin.GallerytypeHandle{}, "GET:Index")
	beego.Router("/admin/gallerytype/add", &admin.GallerytypeHandle{}, "GET:ToAdd")

	beego.Router("/admin/gallery", &admin.GalleryHandle{}, "GET:Index")
	beego.Router("/admin/gallery/add", &admin.GalleryHandle{}, "GET:ToAdd")

	beego.Router("/admin/articletype", &admin.ArticletypeHandle{}, "GET:Index")
	beego.Router("/admin/articletype/add", &admin.ArticletypeHandle{}, "GET:ToAdd")

	beego.Router("/admin/article", &admin.ArticleHandle{}, "GET:Index")
	beego.Router("/admin/article/add", &admin.ArticleHandle{}, "GET:ToAdd")

	beego.Router("/admin/comment", &admin.CommentHandle{}, "GET:Index")

	beego.Router("/admin/manager", &admin.ManagerHandle{}, "GET:Index")
	beego.Router("/admin/manager/update", &admin.ManagerHandle{}, "GET:Update")
	beego.Router("/admin/manager/add", &admin.ManagerHandle{}, "GET:ToAdd")

	beego.Router("/admin/log", &admin.RecordHandle{}, "GET:Index")

	beego.Router("/admin/link", &admin.LinkHandle{}, "GET:Index")
	beego.Router("/admin/link/add", &admin.LinkHandle{}, "GET:ToAdd")
}
