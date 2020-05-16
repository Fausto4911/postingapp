package model

import "time"

type Post struct {
	Id       uint16    `json:"id"`
	User     uint16    `json:"user"`
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Category uint16    `json:"category"`
	Creation time.Time `json:"creation"`
	Comments []string  `json:"comments"`
	Image    string    `json:"image"`
}
