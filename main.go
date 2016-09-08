package main

import (
	_ "easyui_crud/models"
	"github.com/astaxie/beego"
	_ "personal_website/routers"
)

const VERSION = "1.0.0"

func main() {
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}
