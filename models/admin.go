package models

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
}

func (admin *Admin) TableName() string {
	return "admin"
}
