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
	Store(comment Comment)
	FindById(id int64)
}
