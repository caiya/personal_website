// contact
package controllers

type ContactController struct {
	BaseHandle
}

func (this *ContactController) Index() {
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_联系")
	this.RenderHtml("layout.html", "contact.html", datas)
}
