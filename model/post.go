package model

import (
	"errors"
	"time"
)

type Post struct {
	PostId       int64     `json:"post_id"`
	Client       Client    `json:"client"`
	Header       string    `json:"header"`
	Votes        int64     `json:"votes"`
	Category     string    `json:"category"`
	CreationDate time.Time `json:"creation_date"`
	Comments     []Comment `json:"comments"`
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
