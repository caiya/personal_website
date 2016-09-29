package admin

import (
	"personal_website/controllers"
	"personal_website/models"
)

type ArticletypeHandle struct {
	controllers.BaseHandle
}

//照片分类首页
func (this *ArticletypeHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/article_category.html", "c5")
	return
}

//新增页
func (this *ArticletypeHandle) ToAdd() {
	articletype := &models.Articletype{}
	articletypes := articletype.GetAll()
	this.Data["articletypes"] = articletypes
	this.RspTemp("admin/layout.html", "admin/add_article_category.html", "c5")
	return
}
