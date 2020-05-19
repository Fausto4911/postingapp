package model

import "time"

type Comment struct {
	Id       uint16    `json:"id"`
	User     AppUser      `json:"user"`
	Text     string    `json:"text"`
	CreateAt time.Time `json:"create_at"`
	Post     Post      `json:"post"`
	Parent   *Comment  `json:"parent"`
	Child    *Comment  `json:""child`
}
