package main

import (
	_ "easyui_crud/models"
	_ "personal_website/routers"

	"github.com/astaxie/beego"
)

const VERSION = "1.0.0"

func main() {
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}
