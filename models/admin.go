package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name" valid:"Required;MaxSize(50)"`
	Nickname  string `json:"nickname" form:"nickname" valid:"Required;MaxSize(50)"`
	Pass      string `json:"pass" form:"pass" valid:"MaxSize(50)"`
	Repass    string `form:"repass" orm:"-" valid:"MaxSize(50)"` //再次输入的密码，仅仅用于修改密码时校验
	Email     string `json:"email" form:"email" valid:"Required;Email"`
	Mobile    string `json:"mobile" form:"mobile" valid:"Required;Mobile"`
	Address   string `json:"address" form:"address" valid:"Required;MaxSize(50)"`
	Addtime   int    `json:"addtime"`
	Lastlogin int    `json:"lastlogin"`
	Lastip    string `json:"lastip"`
}

func (admin *Admin) TableName() string {
	return "admin"
}

func (admin *Admin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(admin)
}

func init() {
	orm.RegisterModel(new(Admin)) //注册数据库模型
}

//新增
func (admin *Admin) Insert() error {
	if _, err := orm.NewOrm().Insert(admin); err != nil {
		return err
	}
	return nil
}

//删除
func (admin *Admin) Delete() error {
	if _, err := orm.NewOrm().Delete(admin); err != nil {
		return err
	}
	return nil
}

//修改
func (admin *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(admin, fields...); err != nil {
		return err
	}
	return nil
}

//查询
func (admin *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(admin, fields...); err != nil {
		return err
	}
	return nil
}

//分页查询
func (admin *Admin) GetList(query map[string]string, fields []string, sortby []string, order []string, page int64, pageSize int64) (ml []interface{}, count int64, err error) {
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	var ad Admin
	qs := ad.Query()
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1) //-1代表全部替换
		qs = qs.Filter(k, v)
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			orderby := ""
			for i, v := range sortby {
				if strings.ToLower(order[i]) == "asc" {
					orderby = v
				} else if strings.ToLower(order[i]) == "desc" {
					orderby = "-" + v
				} else {
					return nil, 0, errors.New("查询参数有误")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) == 1 { //单独一个asc或者dsc
			orderby := ""
			for i, v := range sortby {
				if strings.ToLower(order[i]) == "asc" {
					orderby = v
				} else if strings.ToLower(order[i]) == "desc" {
					orderby = "-" + v
				} else {
					return nil, 0, errors.New("查询参数有误")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("查询参数有误")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("查询参数有误")
		}
	}
	var admins []Admin
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(pageSize, offset).All(&admins, fields...); err == nil { //成功
		if len(fields) == 0 { //查询全部字段
			for _, v := range admins {
				ml = append(ml, v)
			}
		} else { //根据fields字段进行匹配输出
			var m map[string]interface{}
			for _, v := range admins {
				m = make(map[string]interface{}, 0)
				val := reflect.ValueOf(v)
				for _, field := range fields {
					m[field] = val.FieldByName(field).Interface()
				}
				ml = append(ml, m)
			}
		}
		count, _ := qs.Count()
		return ml, count, nil
	}
	return nil, 0, err
}
