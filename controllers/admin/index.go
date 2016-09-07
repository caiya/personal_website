package admin

import (
	"fmt"
)

type IndexHandle struct {
	baseHandle
}

func (this *IndexHandle) Index() {
	this.TplName = "admin/login.html"
}

func (this *IndexHandle) Login() {
	fmt.Println("#############")
	this.Ctx.WriteString("欢迎登录成功")
}
