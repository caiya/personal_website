package utils

import (
	"html/template"
	"time"

	"github.com/astaxie/beego"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//加added
func Up(in int, uped int) (out int) {
	return in + uped
}

//减added
func Down(in int, downed int) (out int) {
	return in - downed
}

//格式化时间戳
func FormatTimeStamp(timestamp int, layout string) (out string) {
	out = time.Unix(int64(timestamp), 0).Format(layout)
	return
}

//输出market文本为html
func ToHtmlFromMarkdown(s string) template.HTML {
	unsafe := blackfriday.MarkdownCommon([]byte(s))
	s1 := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return template.HTML(s1)
}

func init() {
	beego.AddFuncMap("Up", Up)
	beego.AddFuncMap("Down", Down)
	beego.AddFuncMap("FormatTimeStamp", FormatTimeStamp)
	beego.AddFuncMap("ToHtmlFromMarkdown", ToHtmlFromMarkdown)
}
