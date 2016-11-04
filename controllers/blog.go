// blog
package controllers

type BlogController struct {
	BaseHandle
}

func (this *BlogController) Index() {
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_博客")
	this.RenderHtml("layout.html", "blog.html", datas)
}
