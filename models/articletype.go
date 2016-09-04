package models

type Articletype struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Orderno int    `json:"orderno"`
}

func (articletype *Articletype) TableName() string {
	return "articletype"
}
