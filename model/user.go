package model

import "time"

type User struct {
	Id       uint16    `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Image    string    `json:"image"`
	Date     time.Time `json:date`
}
