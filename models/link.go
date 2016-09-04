package models

type Link struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Orderno int    `json:"orderno"`
}

func (link *Link) TableName() string {
	return "link"
}
