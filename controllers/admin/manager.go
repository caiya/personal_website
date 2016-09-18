package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
)

type ManagerHandle struct {
	controllers.BaseHandle
}

//后台用户首页
func (this *ManagerHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/manager.html", "c8")
	return
}

//新增页
func (this *ManagerHandle) ToAdd() {
	this.RspTemp("admin/layout.html", "admin/add_manager.html", "c8")
	return
}

//修改页
func (this *ManagerHandle) Update() {

}

//分页查询
func (this *ManagerHandle) GetList() {
	admin := &models.Admin{}
	admins, count, err := admin.GetList(nil, nil, nil, nil, 2, 2)

	rsp := make(map[string]interface{})

	if err == nil {
		rsp["datas"] = admins
		rsp["count"] = count
		this.Data["json"] = rsp
		this.ServeJSON()
	} else {
		fmt.Println("报错了。。。。")
	}
}
