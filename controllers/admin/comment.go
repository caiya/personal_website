package admin

import (
	"personal_website/controllers"
)

type CommentHandle struct {
	controllers.BaseHandle
}

//评论首页
func (this *CommentHandle) Index() {
	this.RspTemp("admin/layout.html", "admin/comment.html", "c7")
	return
}
