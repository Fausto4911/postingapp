package postgres

import (
	"database/sql"
)

type postgresPostRepository struct {
	db *sql.DB
}

/*
    Store(post Post) (Post, error)
	FindById(id int64) (Post, error)
	FindAll() ([]Post, error)
	Update(client Client) (Client, error)
	Delete(client Client) (bool, error)
*/
//
//func (p *postgresPostRepository) Store(ctx context.Context, post model.Post) (stored model.Post, err error) {
//	result, err := p.db.ExecContext(ctx, "INSERT INTO post (post_id, user, header, votes, category, creation_date) post VALUES(nextval('post_seq'), $1, $2, $3, $4, $5)",
//		post.Id,
//		post.Title,
//		post.Text,
//		post.Category,
//		time.Now(),
//	)
//	if err != nil {
//		return model.Post{}, errors.New("cannot store post")
//	}
//	//lastId, er := result.LastInsertId()
//	//if er != nil {
//	//	return model.Post{}, errors.New("cannot retrieve last inserted id")
//	//}
//	//post.Id = lastId
//	return post, nil
//}
