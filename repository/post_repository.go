package repository

import (
	"go-gin-blog-api/model"
	"sync"
)

type PostRepository interface {
	Save(post model.Post) model.Post
	FindAll() []model.Post
}

type postRepository struct {
	mu    sync.Mutex
	posts []model.Post
}

func NewPostRepository() PostRepository {
	return &postRepository{
		posts: make([]model.Post, 0),
	}
}

func (r *postRepository) Save(post model.Post) model.Post {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.posts = append(r.posts, post)
	return post
}

func (r *postRepository) FindAll() []model.Post {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]model.Post(nil), r.posts...)
}
