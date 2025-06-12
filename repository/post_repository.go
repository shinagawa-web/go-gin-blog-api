package repository

import (
	"go-gin-blog-api/model"
	"sync"
)

type PostRepository interface {
	Save(post model.Post) model.Post
	FindAll() []model.Post
	FindByID(id string) (*model.Post, bool)
	Update(id string, updated model.Post) (*model.Post, bool)
	Delete(id string) bool
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

func (r *postRepository) FindByID(id string) (*model.Post, bool) {
	for _, post := range r.posts {
		if post.ID == id {
			return &post, true
		}
	}
	return nil, false
}

func (r *postRepository) Update(id string, updated model.Post) (*model.Post, bool) {
	for i, post := range r.posts {
		if post.ID == id {
			if updated.Title != "" {
				r.posts[i].Title = updated.Title
			}
			if updated.Content != "" {
				r.posts[i].Content = updated.Content
			}
			if updated.Author != "" {
				r.posts[i].Author = updated.Author
			}
			return &r.posts[i], true
		}
	}
	return nil, false
}

func (r *postRepository) Delete(id string) bool {
	for i, post := range r.posts {
		if post.ID == id {
			r.posts = append(r.posts[:i], r.posts[i+1:]...)
			return true
		}
	}
	return false
}
