/**
前台首页
*/
package controllers

type IndexController struct {
	BaseHandle
}

func (this *IndexController) Index() {
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_首页")
	this.RenderHtml("layout.html", "index.html", datas)
}
