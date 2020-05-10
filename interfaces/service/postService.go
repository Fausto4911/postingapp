package interfaces

import (
	"context"
	"database/sql"
	"postingapp/domain"
	"time"
)

// Store(post Post) (Post, error)
// FindById(id int64) (Post, error)
// FindAll() ([]Post, error)
// Update(client Client) (Client, error)
// Delete(client Client) (bool, error)

func (service *PostService) Store(post domain.Post) (domain.Post, error) {
	result, err := service.db.ExecContext(service.ctx, "INSERT INTO(post_id, user, header, votes, category, creation_date) post VALUES(nextval('post_seq'), $1, $2, $3, $4, $5)",
		post.Client.UserId,
		post.Header,
		post.Votes,
		post.Category,
		time.Now()
	)
}

type PostService struct {
	Repo PostRepository
	db   *sql.DB
	ctx  context.Context
}
