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

	//系统设置
	beego.Router("/admin/setting", &admin.SettingHandle{}, "GET:Index")
	beego.Router("/admin/setting/update", &admin.SettingHandle{}, "POST:Update")

	//单页模板
	beego.Router("/admin/template", &admin.TemplateHandle{}, "GET:Index")
	beego.Router("/admin/template/add", &admin.TemplateHandle{}, "GET:ToAdd")
	beego.Router("/admin/template/add", &admin.TemplateHandle{}, "POST:Add")
	beego.Router("/admin/template/update/?:id", &admin.TemplateHandle{}, "GET:ToUpdate")
	beego.Router("/admin/template/delete/?:id", &admin.TemplateHandle{}, "GET:Delete")

	//相册分类
	beego.Router("/admin/gallerytype", &admin.GallerytypeHandle{}, "GET:Index")
	beego.Router("/admin/gallerytype/add", &admin.GallerytypeHandle{}, "GET:ToAdd")

	//相册
	beego.Router("/admin/gallery", &admin.GalleryHandle{}, "GET:Index")
	beego.Router("/admin/gallery/add", &admin.GalleryHandle{}, "GET:ToAdd")

	//博客类别
	beego.Router("/admin/articletype", &admin.ArticletypeHandle{}, "GET:Index")
	beego.Router("/admin/articletype/add", &admin.ArticletypeHandle{}, "GET:ToAdd")

	//文章|博客
	beego.Router("/admin/article/?:page", &admin.ArticleHandle{}, "*:Index")
	beego.Router("/admin/article/add", &admin.ArticleHandle{}, "GET:ToAdd")
	beego.Router("/admin/article/add", &admin.ArticleHandle{}, "POST:Add")
	beego.Router("/admin/article/update", &admin.ArticleHandle{}, "GET:ToUpdate")
	beego.Router("/admin/article/update", &admin.ArticleHandle{}, "POST:Update")
	beego.Router("/admin/article/delete/:id", &admin.ArticleHandle{}, "GET:Delete")

	//评论
	beego.Router("/admin/comment", &admin.CommentHandle{}, "GET:Index")
	beego.Router("/admin/comment/delete/?:id", &admin.CommentHandle{}, "GET:Delete")

	//后台管理员
	beego.Router("/admin/manager/?:page", &admin.ManagerHandle{}, "GET:Index")
	beego.Router("/admin/manager/update/?:id", &admin.ManagerHandle{}, "GET:ToUpdate")
	beego.Router("/admin/manager/add", &admin.ManagerHandle{}, "GET:ToAdd")
	beego.Router("/admin/manager/add", &admin.ManagerHandle{}, "POST:Add")
	beego.Router("/admin/manager/delete/?:id", &admin.ManagerHandle{}, "GET:Delete")

	//日志模块
	beego.Router("/admin/log/?:page", &admin.RecordHandle{}, "*:Index")

	//链接模块
	beego.Router("/admin/link/delete/?:id", &admin.LinkHandle{}, "GET:Delete") //单个删除
	beego.Router("/admin/link/delete", &admin.LinkHandle{}, "POST:Del")        //批量删除
	beego.Router("/admin/link/?:page", &admin.LinkHandle{}, "*:Index")
	beego.Router("/admin/link/add", &admin.LinkHandle{}, "GET:ToAdd")
	beego.Router("/admin/link/add", &admin.LinkHandle{}, "POST:Add")
	beego.Router("/admin/link/update/?:id", &admin.LinkHandle{}, "GET:ToUpdate")

	//上传
	beego.Router("/upload", &admin.IndexHandle{}, "POST:Upload")

}
