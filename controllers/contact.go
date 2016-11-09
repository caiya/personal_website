// contact
package controllers

import (
	"personal_website/models"
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
