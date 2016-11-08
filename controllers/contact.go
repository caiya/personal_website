// contact
package controllers

import (
	"personal_website/models"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
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
	unsafe := blackfriday.MarkdownCommon([]byte(s))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	this.Data["temp"] = unsafe
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_联系")
	this.RenderHtml("layout.html", "contact.html", datas)
}
