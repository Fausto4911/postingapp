package domain

import "time"

type Comment struct {
	CommentId    int64
	User         Client
	Description  string
	CreationDate time.Time
	Post         Post
}

type CommentRepository interface {
	Store(comment Comment) (Comment, error)
	FindById(id int64) (Comment, error)
	FindAll() ([]Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(client Client) (bool, error)
}
