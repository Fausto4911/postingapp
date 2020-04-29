package domain

import "time"

type Post struct {
	PostId       int64
	Client       Client
	Header       string
	Votes        int64
	Category     string
	CreationDate time.Time
}

type PostRepository interface {
	Store(post Post) Post
	FindById(id int64) Post
}
