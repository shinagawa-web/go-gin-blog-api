package service

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/repository"
)

type PostService interface {
	Create(post model.Post) model.Post
	List() []model.Post
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{repo: r}
}

func (s *postService) Create(post model.Post) model.Post {
	return s.repo.Save(post)
}

func (s *postService) List() []model.Post {
	return s.repo.FindAll()
}
