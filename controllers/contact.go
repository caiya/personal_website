// contact
package controllers

import (
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (this *ContactController) Index() {
	this.TplName = "contact.html"
}
