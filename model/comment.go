package model

import "time"

type Comment struct {
	Id       uint16    `json:"id"`
	User     uint16    `json:"user"`
	Text     string    `json:"text"`
	Creation time.Time `json:"creation"`
	Post     uint16    `json:"post"`
}
