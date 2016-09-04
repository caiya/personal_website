package models

type Contact struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Message string `json:"message"`
}

func (contact *Contact) TableName() string {
	return "contact"
}
