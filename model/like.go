package model

import "time"

type Like struct {
	Id       uint16    `json:"id"`
	User     AppUser      `json:"post"`
	CreateAt time.Time `json:"create_at"`
}
