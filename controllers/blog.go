// blog
package controllers

import (
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Index() {
	this.TplName = "blog.html"
}
