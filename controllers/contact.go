// contact
package controllers

import (
	"fmt"
	"personal_website/models"
	"time"
)

type ContactController struct {
	BaseHandle
}

func (this *ContactController) Index() {
	var temp = new(models.Template)
	temp.Name = "contact"
	err := temp.Read("name")
	var s string
	if err != nil {
		s = ""
	} else {
		s = temp.Content
	}
	this.Data["temp"] = s
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_联系")
	this.RenderHtml("layout.html", "contact.html", datas)
}

//联系我
func (this *ContactController) Contact() {
	contact := models.Contact{}
	if err := this.ParseForm(&contact); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v", contact)
	if contact.Id == 0 { //新增
		contact.Ip = this.GetIP()
		contact.Addtime = int(time.Now().Unix())
		contact.Insert()
	} else { //修改
		contact.Update()
	}
	this.Redirect("/contact", 302)
}
