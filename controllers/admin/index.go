package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"personal_website/utils"
)

type IndexHandle struct {
	controllers.BaseHandle
}

func (this *IndexHandle) Index() {
	this.TplName = "admin/login.html"
}

func (this *IndexHandle) Login() {
	username := this.GetString("userName")
	password := this.GetString("password")
	u := &models.Admin{Name: username}
	u.Read("Name")
	if u.Id == 0 {
		this.Data["errmsg"] = "当前用户名不存在"
		this.Data["username"] = username
		this.Data["password"] = password
		this.TplName = "admin/login.html"
		return
	} else if u.Pass != utils.Md5(password) {
		this.Data["errmsg"] = "密码错误，请重输"
		this.Data["username"] = username
		this.Data["password"] = password
		this.TplName = "admin/login.html"
		return
	}
	this.Redirect("/admin/main", 302)
}

func (this *IndexHandle) Main() {
	this.TplName = "admin/index.html"
}
