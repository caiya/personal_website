package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int          `json:"id"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Addtime  int          `json:"addtime"`
	Uptime   int          `json:"uptime"`
	User     *User        `json:"user" orm:"rel(fk)"`
	Link     string       `json:"link"`
	Intro    string       `json:"intro"`
	Type     *Articletype `json:"type" orm:"rel(fk)"`
	Comments []*Comment   `orm:"reverse(many)"` //反向一对多关联
}

func (article *Article) TableName() string {
	return "article"
}

func (article *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(article)
}

//注册模型
func init() {
	orm.RegisterModel(new(Article))
}

//新增
func (article *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(article); err != nil {
		return err
	}
	return nil
}

//删除
func (article *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(article); err != nil {
		return err
	}
	return nil
}

//修改
func (article *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(article, fields...); err != nil {
		return err
	}
	return nil
}

//查询
func (article *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(article, fields...); err != nil {
		return err
	}
	return nil
}

//分页查询
func (article *Article) GetList(query map[string]string, fields []string, sortby []string, order []string, page int64, pageSize int64) (ml []interface{}, count int64, err error) {
	var offset int64
	if page > 1 {
		offset = (page - 1) * pageSize
	}
	var ad Article
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
	var articles []Article
	//将切片打散传入
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(pageSize, offset).RelatedSel().All(&articles, fields...); err == nil { //成功
		if len(fields) == 0 { //查询全部字段
			for _, v := range articles {
				ml = append(ml, v)
			}
		} else { //根据fields字段进行匹配输出
			var m map[string]interface{}
			for _, v := range articles {
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
