package domain

import (
	"errors"
	"time"
)

type Post struct {
	PostId       int64
	Client       Client
	Header       string
	Votes        int64
	Category     string
	CreationDate time.Time
	Comments     []Comment
}

type PostRepository interface {
	Store(post Post) (Post, error)
	FindById(id int64) (Post, error)
	FindAll() ([]Post, error)
	Update(client Client) (Client, error)
	Delete(client Client) (bool, error)
}

func (post *Post) AddComment(comment Comment) error {
	post.Comments = append(post.Comments, comment)
	return nil
}

func (post *Post) AddVote(vote int64) error {
	post.Votes = post.Votes + vote
	return nil
}

func (post *Post) addCategory(category string) error {
	if len(category) < 1 {
		return errors.New("Cannot save the category of lenght less than 1")
	}
	post.Category = category
	return nil
}
