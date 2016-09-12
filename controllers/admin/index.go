package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"personal_website/utils"
)

type IndexHandle struct {
	controllers.BaseHandle
}

type MainHandle struct {
	controllers.BaseHandle
}

//首页跳转
func (this *IndexHandle) Index() {
	this.TplName = "admin/login.html"
}

//登录
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
	u.Pass = ""
	this.SetSession("currUser", u) //设置session
	this.Redirect("/admin/main", 302)
}

//退出登录
func (this *IndexHandle) Exit() {
	this.DelSession("currUser")
	this.Redirect("/admin", 302)
}

//主页跳转
func (this *MainHandle) Main() {
	u := this.GetSession("currUser")
	if u != nil {
		this.Data["currUser"] = u
		this.TplName = "admin/index.html"
	} else {
		this.Redirect("/admin", 302)
	}
}
