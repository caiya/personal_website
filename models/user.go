package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Mobile    string    `json:"mobile"`
	Age       int       `json:"age"`
	Sex       int       `json:"sex"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Pass      string    `json:"pass"`
	Addtime   time.Time `json:"addtime"`
	Lastlogin time.Time `json:"lastlogin"`
}

func (user *User) TableName() string {
	return "user"
}