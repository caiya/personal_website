package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"personal_website/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/validation"
)

const DEFAULTPASS string = "123456"

type ManagerHandle struct {
	controllers.BaseHandle
}

//后台用户首页
func (this *ManagerHandle) Index() {
	admin := &models.Admin{}
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}

	admins, count, err := admin.GetList(nil, nil, []string{"id"}, []string{"desc"}, int64(page), PAGESIZE)
	if err == nil {
		this.Data["datas"] = admins //记录数
		this.Data["count"] = count  //总记录数
	} else {
		this.Data["datas"] = nil
		this.Data["count"] = 0
	}
	var pageCount int64
	if count%PAGESIZE == 0 {
		pageCount = count / PAGESIZE
	} else {
		pageCount = count/PAGESIZE + 1
	}
	this.Data["pageCount"] = pageCount //总页数
	this.Data["page"] = page           //当前页
	this.RspTemp("admin/layout.html", "admin/manager.html", "c8")
	return
}

//新增页
func (this *ManagerHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
	return
}

//修改页
func (this *ManagerHandle) ToUpdate() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	admin := &models.Admin{Id: id}
	admin.Read()
	this.Data["admin"] = admin
	this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
	return
}

//新增或者修改操作
func (this *ManagerHandle) Add() {
	admin := &models.Admin{}
	if err := this.ParseForm(admin); err != nil {
		admin.Read("id")
		this.Data["errmsg"] = err.Error()
		this.Data["admin"] = admin
		this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(admin)
	if err != nil { //程序出错
		this.Data["errmsg"] = err.Error()
		admin.Read("id")
		this.Data["admin"] = admin
		this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
		return
	}
	if !b { //校验失败
		for _, err := range valid.Errors {
			admin.Read("id")
			this.Data["admin"] = admin
			this.Data["errmsg"] = err.Error()
			this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
			return
		}
	}
	if admin.Id == 0 { //新增
		if admin.Pass == "" {
			admin.Pass = DEFAULTPASS
		}
		admin.Pass = utils.Md5(admin.Pass)
		admin.Addtime = int(time.Now().Unix())
		admin.Insert()
	} else { //修改
		if admin.Pass != "" && admin.Repass != "" {
			admin.Pass = utils.Md5(admin.Pass)
			admin.Update("id", "name", "nickname", "address", "mobile", "email", "pass")
		} else {
			admin.Update("id", "name", "nickname", "address", "mobile", "email")
		}
	}
	//成功后直接重定向
	this.Redirect("/admin/manager", 302)
}

//删除
func (this *ManagerHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	admin := &models.Admin{Id: id}
	admin.Delete()
	this.Redirect("/admin/manager", 302)
}
