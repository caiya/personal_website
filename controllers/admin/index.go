package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"personal_website/utils"
	"strings"
	"time"
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
	username := strings.TrimSpace(this.GetString("userName"))
	password := strings.TrimSpace(this.GetString("password"))
	u := &models.Admin{Name: username}
	u.Read("name")
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

	u.Lastlogin = int(time.Now().Unix())
	u.Lastip = this.GetIP()
	u.Update("lastlogin", "lastip") //更新登录日志

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
		this.Layout = "admin/layout.html"
		this.TplName = "admin/index.html"
	} else {
		this.Redirect("/admin", 302)
	}
}

//上传图片
func (this *IndexHandle) Upload() {
	this.UploadImg()
	return
}
