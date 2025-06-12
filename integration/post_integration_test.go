package integration_test

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_PostServiceLifecycle(t *testing.T) {
	repo := repository.NewPostRepository()
	svc := service.NewPostService(repo)

	// Create
	post := model.Post{
		ID:      "200",
		Title:   "Test Lifecycle",
		Content: "This post will be updated and deleted.",
		Author:  "TestBot",
	}
	created := svc.Create(post)

	// GetByID
	got, found := svc.GetByID("200")
	assert.True(t, found)
	assert.Equal(t, created, got)

	// Update
	updatedPost := model.Post{
		Title:   "Updated Title",
		Content: "Updated content",
	}
	updated, ok := svc.Update("200", updatedPost)
	assert.True(t, ok)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "Updated content", updated.Content)

	// List
	all := svc.List()
	assert.Len(t, all, 1)
	assert.Equal(t, "Updated Title", all[0].Title)

	// Delete
	deleted := svc.Delete("200")
	assert.True(t, deleted)

	// Confirm Deletion
	_, found = svc.GetByID("200")
	assert.False(t, found)
}
