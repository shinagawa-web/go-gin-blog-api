package repository_test

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndFindByID(t *testing.T) {
	repo := repository.NewPostRepository()
	post := model.Post{ID: "1", Title: "First", Content: "Hello", Author: "Alice"}

	repo.Save(post)
	found, ok := repo.FindByID("1")

	assert.True(t, ok)
	assert.Equal(t, "First", found.Title)
	assert.Equal(t, "Alice", found.Author)
}

func TestFindAll(t *testing.T) {
	repo := repository.NewPostRepository()
	repo.Save(model.Post{ID: "1"})
	repo.Save(model.Post{ID: "2"})

	all := repo.FindAll()
	assert.Len(t, all, 2)
}

func TestUpdate_Success(t *testing.T) {
	repo := repository.NewPostRepository()
	repo.Save(model.Post{ID: "1", Title: "Old"})

	updated := model.Post{Title: "New"}
	post, ok := repo.Update("1", updated)

	assert.True(t, ok)
	assert.Equal(t, "New", post.Title)
}

func TestUpdate_Failure(t *testing.T) {
	repo := repository.NewPostRepository()
	updated := model.Post{Title: "New"}
	post, ok := repo.Update("99", updated)

	assert.False(t, ok)
	assert.Nil(t, post)
}

func TestDelete_Success(t *testing.T) {
	repo := repository.NewPostRepository()
	repo.Save(model.Post{ID: "1"})

	ok := repo.Delete("1")
	assert.True(t, ok)

	_, found := repo.FindByID("1")
	assert.False(t, found)
}

func TestDelete_Failure(t *testing.T) {
	repo := repository.NewPostRepository()

	ok := repo.Delete("999")
	assert.False(t, ok)
}
