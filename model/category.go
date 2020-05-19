package model

import "time"

type Category struct {
	Id       uint16    `json:"id"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
}

type CategoryRepository interface {
	Store(category Category) error
	GetAll() ([]Category, error)
	GetById(id int) (Category, error)
	Delete(id int) error
	Update(category Category) error
}
