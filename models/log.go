package models

import (
	"strings"

	"errors"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id     int    `json:"id"`
	Optime int    `json:"optime"`
	Opname string `json:"opname"`
	Record string `json:"record"`
	Ipaddr string `json:"ipaddr"`
}

func (record *Record) TableName() string {
	return "log"
}

func (record *Record) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(record)
}

//注册模型
func init() {
	orm.RegisterModel(new(Record))
}

//新增
func (record *Record) Insert() error {
	if _, err := orm.NewOrm().Insert(record); err != nil {
		return err
	}
	return nil
}

//删除
func (record *Record) Delete() error {
	if _, err := orm.NewOrm().Delete(record); err != nil {
		return err
	}
	return nil
}

//修改
func (record *Record) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(record, fields...); err != nil {
		return err
	}
	return nil
}

//查询
func (record *Record) Read(fields ...string) error {
	if err := orm.NewOrm().Read(record, fields...); err != nil {
		return err
	}
	return nil
}

//分页查询
func (record *Record) GetList(query map[string]string, fields []string, sortby []string, order []string, page int64, pageSize int64) (ml []interface{}, count int64, err error) {
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	var ad Record
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
	var records []Record
	//将切片打散传入
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(pageSize, offset).All(&records, fields...); err == nil { //成功
		if len(fields) == 0 { //查询全部字段
			for _, v := range records {
				ml = append(ml, v)
			}
		} else { //根据fields字段进行匹配输出
			var m map[string]interface{}
			for _, v := range records {
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
