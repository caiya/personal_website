package models

type Gallery struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Imgurl  string `json:"imgurl"`
	Orderno int    `json:"orderno"`
}

func (gallery *Gallery) TableName() string {
	return "gallery"
}
