package admin

import (
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type RecordHandle struct {
	controllers.BaseHandle
}

//日志首页
func (this *RecordHandle) Index() {
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	record := &models.Record{}

	sortby := []string{"id", "optime"} //id、optime逆序
	order := []string{"desc", "desc"}

	records, count, err := record.GetList(nil, nil, sortby, order, int64(page), PAGESIZE)
	if err == nil {
		this.Data["datas"] = records //记录数
		this.Data["count"] = count   //总记录数
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
	this.RspTemp("admin/layout.html", "admin/manager_log.html", "c9")
	return
}
