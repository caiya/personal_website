package models

type Gallerytype struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Orderno int    `json:"orderno"`
}

func (gallerytype *Gallerytype) TableName() string {
	return "gallerytype"
}
