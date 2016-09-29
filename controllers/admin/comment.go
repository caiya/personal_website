package admin

import (
	"fmt"
	"personal_website/controllers"
	"personal_website/models"
	"strconv"
)

type CommentHandle struct {
	controllers.BaseHandle
}

//评论首页
func (this *CommentHandle) Index() {
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	if page == 0 {
		page = 1
	}
	comment := &models.Comment{}

	sortby := []string{"id"} //id、optime逆序
	order := []string{"desc"}

	comments, count, err := comment.GetList(nil, nil, sortby, order, int64(page), PAGESIZE)
	if err == nil {
		this.Data["datas"] = comments //记录数
		this.Data["count"] = count    //总记录数
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

	fmt.Println(comments)

	this.Data["pageCount"] = pageCount //总页数
	this.Data["page"] = page           //当前页
	this.RspTemp("admin/layout.html", "admin/comment.html", "c7")
	return
}

//删除
func (this *CommentHandle) Delete() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	comment := &models.Comment{Id: id}
	comment.Delete()
	this.Redirect("/admin/comment", 302)
}
