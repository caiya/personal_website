package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Contact struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Ip      string `json:"ip" form:"-" orm:"ip"`
	Message string `json:"message" form:"message"`
	Addtime int    `json:"addtime" form:"-"`
}

func (contact *Contact) TableName() string {
	return "contact"
}

func (contact *Contact) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(contact)
}

//注册模型
func init() {
	orm.RegisterModel(new(Contact))
}

//新增
func (contact *Contact) Insert() error {
	if _, err := orm.NewOrm().Insert(contact); err != nil {
		return err
	}
	return nil
}

//单个删除
func (contact *Contact) Delete() error {
	if _, err := orm.NewOrm().Delete(contact); err != nil {
		return err
	}
	return nil
}

//批量删除
func (contact *Contact) Del(idstr string) error {
	ids := strings.Split(idstr, ",")
	num, err := contact.Query().Filter("id__in", ids).Delete()
	if num > 0 && err == nil {
		return nil
	} else {
		return err
	}
}

//修改
func (contact *Contact) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(contact, fields...); err != nil {
		return err
	}
	return nil
}

//查询
func (contact *Contact) Read(fields ...string) error {
	if err := orm.NewOrm().Read(contact, fields...); err != nil {
		return err
	}
	return nil
}

//列表查询
func (contact *Contact) GetList(query map[string]string, fields []string, sortby []string, order []string, page int64, pageSize int64) (ml []interface{}, count int64, err error) {
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	var ad Contact
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
	var contacts []Contact
	//将切片打散传入
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(pageSize, offset).All(&contacts, fields...); err == nil { //成功
		if len(fields) == 0 { //查询全部字段
			for _, v := range contacts {
				ml = append(ml, v)
			}
		} else { //根据fields字段进行匹配输出
			var m map[string]interface{}
			for _, v := range contacts {
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
