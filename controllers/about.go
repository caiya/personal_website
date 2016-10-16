// about
package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index() {
	this.TplName = "about.html"
}
