package main

import (
	_ "easyui_crud/models"
	"html/template"
	"net/http"
	_ "personal_website/routers"

	_ "personal_website/utils"

	"github.com/astaxie/beego"
)

const VERSION = "1.0.0"

func main() {

	beego.AppConfig.Set("version", VERSION)
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}

//404处理
func page_not_found(rw http.ResponseWriter, rq *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.AppPath + "/views/404.html")
	t.Execute(rw, nil)
}
