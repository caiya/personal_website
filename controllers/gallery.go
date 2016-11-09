package controllers

type GalleryController struct {
	BaseHandle
}

func (this *GalleryController) Index() {
	renderDatas := make([]string, 0)
	datas := append(renderDatas, "title_首页")
	this.RenderHtml("layout.html", "gallery.html", datas)
}
