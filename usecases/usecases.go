package usecases

import (
	"errors"
	"postingapp/domain"
	"strings"
	"time"
)

type PostInteractor struct {
	CategoryRepository domain.CategoryRepository
	ClientRepository   domain.ClientRepository
	PostRepository     domain.PostRepository
	CommentRepository  domain.CommentRepository
}

// Cliente hace un comentario a un post
func (postI *PostInteractor) AddComment(userId, postId int64, comment string) error {
	if len(comment) < 1 {
		return errors.New("Cannot save null comments")
	}
	client := postI.ClientRepository.FindById(userId)
	post := postI.PostRepository.FindById(postId)
	var comm domain.Comment
	comm.User = client
	comm.Post = post
	comm.Description = comment
	comm.CreationDate = time.Now()
	postI.CommentRepository.Store(comm)
	return nil
}

//Cliente da voto a un post
func (postI *PostInteractor) AddVote(postId int64) error {
	post := postI.PostRepository.FindById(postId)
	post.Votes = post.Votes + 1
	postI.PostRepository.Store(post)
	return nil
}

//Cliente asocia una categoria a un post
func (postI *PostInteractor) AddCategory(postId int64, category string) error {
	post := postI.PostRepository.FindById(postId)
	post.Category = category
	postI.PostRepository.Store(post)
	return nil
}

//El Cliente hace login
func LoginUser(userName, passw string, clientRepo domain.ClientRepository) error {
	for _, element := range clientRepo.FindAll() {
		if strings.Compare(element.Password, passw) == 0 && strings.Compare(element.Username, userName) {
			return nil
		}
	}

	return errors.New("Cannot find user in data base")
}
