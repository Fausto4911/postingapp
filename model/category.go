package model

import "time"

type Category struct {
	Id       uint16    `json:"id"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
}
