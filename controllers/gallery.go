package controllers

import (
	"github.com/astaxie/beego"
)

type GalleryController struct {
	beego.Controller
}

func (this *GalleryController) Index() {

	this.TplName = "gallery.html"
}
