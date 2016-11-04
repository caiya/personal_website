// about
package controllers

type AboutController struct {
	BaseHandle
}

func (this *AboutController) Index() {
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_关于我")
	this.RenderHtml("layout.html", "about.html", datas)
}
