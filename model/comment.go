package model

import "time"

type Comment struct {
	CommentId    int64     `json:"comment_id"`
	User         Client    `json:"user"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creation_date"`
	Post         Post      `json:"post"`
}

type CommentRepository interface {
	Store(comment Comment) (Comment, error)
	FindById(id int64) (Comment, error)
	FindAll() ([]Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(client Client) (bool, error)
}
