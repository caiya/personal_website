package admin

import (
	"personal_website/controllers"
)

type RecordHandle struct {
	controllers.BaseHandle
}

//日志首页
func (this *RecordHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/manager_log.html", "c9")
	return
}
