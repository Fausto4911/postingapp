package model

import "time"

type Post struct {
	Id       uint16    `json:"id"`
	User     AppUser      `json:"user"`
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Category Category  `json:"category"`
	CreateAt time.Time `json:"create_at"`
	Image    string    `json:"image"`
}
