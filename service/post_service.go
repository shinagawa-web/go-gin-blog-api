package service

import (
	"go-gin-blog-api/model"
	"sync"
)

type PostService interface {
	Create(post model.Post) model.Post
	List() []model.Post
}

type postService struct {
	mu    sync.Mutex
	posts []model.Post
}

func NewPostService() PostService {
	return &postService{
		posts: make([]model.Post, 0),
	}
}

func (s *postService) Create(post model.Post) model.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.posts = append(s.posts, post)
	return post
}

func (s *postService) List() []model.Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]model.Post(nil), s.posts...)
}
