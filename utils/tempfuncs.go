package utils

import (
	"time"

	"github.com/astaxie/beego"
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

func init() {
	beego.AddFuncMap("Up", Up)
	beego.AddFuncMap("Down", Down)
	beego.AddFuncMap("FormatTimeStamp", FormatTimeStamp)
}
