package service

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/repository"
)

type PostService interface {
	Create(post model.Post) *model.Post
	List() []model.Post
	GetByID(id string) (*model.Post, bool)
	Update(id string, updated model.Post) (*model.Post, bool)
	Delete(id string) bool
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{repo: r}
}

func (s *postService) Create(post model.Post) *model.Post {
	return s.repo.Save(post)
}

func (s *postService) List() []model.Post {
	return s.repo.FindAll()
}

func (s *postService) GetByID(id string) (*model.Post, bool) {
	return s.repo.FindByID(id)
}

func (s *postService) Update(id string, post model.Post) (*model.Post, bool) {
	return s.repo.Update(id, post)
}

func (s *postService) Delete(id string) bool {
	return s.repo.Delete(id)
}
