package models

import (
	"strings"

	"errors"

	"reflect"

	"github.com/astaxie/beego/orm"
)

type Link struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Url     string `json:"url" form:"url"`
	Orderno int    `json:"orderno" form:"orderno"`
}

func (link *Link) TableName() string {
	return "link"
}

func (link *Link) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(link)
}

//注册模型
func init() {
	orm.RegisterModel(new(Link))
}

//新增
func (link *Link) Insert() error {
	if _, err := orm.NewOrm().Insert(link); err != nil {
		return err
	}
	return nil
}

//单个删除
func (link *Link) Delete() error {
	if _, err := orm.NewOrm().Delete(link); err != nil {
		return err
	}
	return nil
}

//批量删除
func (link *Link) Del(idstr string) error {
	ids := strings.Split(idstr, ",")
	num, err := link.Query().Filter("id__in", ids).Delete()
	if num > 0 && err == nil {
		return nil
	} else {
		return err
	}
}

//修改
func (link *Link) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(link, fields...); err != nil {
		return err
	}
	return nil
}

//查询
func (link *Link) Read(fields ...string) error {
	if err := orm.NewOrm().Read(link, fields...); err != nil {
		return err
	}
	return nil
}

//分页查询
func (link *Link) GetList(query map[string]string, fields []string, sortby []string, order []string, page int64, pageSize int64) (ml []interface{}, count int64, err error) {
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	var ad Link
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
	var links []Link
	//将切片打散传入
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(pageSize, offset).All(&links, fields...); err == nil { //成功
		if len(fields) == 0 { //查询全部字段
			for _, v := range links {
				ml = append(ml, v)
			}
		} else { //根据fields字段进行匹配输出
			var m map[string]interface{}
			for _, v := range links {
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
