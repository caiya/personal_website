package utils

import (
	"fmt"

	"github.com/astaxie/beego"
)

func Upload(contro interface{}) {
	this := contro.(beego.Controller)
	f, fh, err := this.GetFile("uploadFile")
	defer f.Close()
	if err != nil {
		fmt.Println("get file error ", err)
		this.Data["json"] = &map[string]interface{}{"path": "", "succ": false}
		this.ServeJSON()
	} else {
		fmt.Println(fh.Filename)
		this.SaveToFile("uploadFile", "static/upload/"+fh.Filename)
		this.Data["json"] = &map[string]interface{}{"path": "/static/upload/" + fh.Filename, "succ": true}
		this.ServeJSON()
	}
}
