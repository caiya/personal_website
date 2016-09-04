package models

import (
	"time"
)

type Comment struct {
	Id      int       `json:"id"`
	Cid     int       `json:"cid"`
	Cname   string    `json:"cname"`
	Cemail  string    `json:"cemail"`
	Content string    `json:"content"`
	Addtime time.Time `json:"addtime"`
	Aid     int       `json:"aid"`
}

func (comment *Comment) TableName() string {
	return "comment"
}
