package model

import "time"

type AppUser struct {
	Id       uint16    `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar_url"`
	CreateAt time.Time `json:"create_at"`
}
