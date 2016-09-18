package utils

import (
	"github.com/astaxie/beego"
)

//加added
func Up(in int, uped int) (out int) {
	out = in + uped
	return
}

//减added
func Down(in int, downed int) (out int) {
	out = in - downed
	return
}

func init() {
	beego.AddFuncMap("Up", Up)
	beego.AddFuncMap("Down", Down)
}
