package models

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.pass")
	timezone := beego.AppConfig.String("db.timezone")

	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn, 5, 30)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

//MD5加密
func Md5Str(s string) (rs string) {
	h := md5.New()
	h.Write([]byte(s))
	rs = hex.EncodeToString(h.Sum(nil))
	return
}
