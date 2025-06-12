package service_test

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックリポジトリ
type mockPostRepository struct {
	mock.Mock
}

func (m *mockPostRepository) FindByID(id string) (*model.Post, bool) {
	args := m.Called(id)
	if post := args.Get(0); post != nil {
		return post.(*model.Post), args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *mockPostRepository) Save(post model.Post) *model.Post {
	args := m.Called(post)
	if p := args.Get(0); p != nil {
		return p.(*model.Post)
	}
	return nil
}

func (m *mockPostRepository) Update(id string, updated model.Post) (*model.Post, bool) {
	args := m.Called(id, updated)
	if p := args.Get(0); p != nil {
		return p.(*model.Post), args.Bool(1)
	}
	return nil, false
}

func (m *mockPostRepository) Delete(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func (m *mockPostRepository) FindAll() []model.Post {
	args := m.Called()
	return args.Get(0).([]model.Post)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(mockPostRepository)
	expected := &model.Post{
		ID: "1", Title: "Gin Guide", Content: "Test content", Author: "Author1",
	}
	mockRepo.On("FindByID", "1").Return(expected, true)

	svc := service.NewPostService(mockRepo)

	post, found := svc.GetByID("1")

	assert.True(t, found)
	assert.Equal(t, expected, post)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := new(mockPostRepository)
	mockRepo.On("FindByID", "999").Return(nil, false)

	svc := service.NewPostService(mockRepo)

	post, found := svc.GetByID("999")

	assert.False(t, found)
	assert.Nil(t, post)
	mockRepo.AssertExpectations(t)
}

func TestCreate_Success(t *testing.T) {
	mockRepo := new(mockPostRepository)
	input := model.Post{Title: "New", Content: "Content", Author: "Dev"}
	expected := &model.Post{ID: "123", Title: "New", Content: "Content", Author: "Dev"}

	mockRepo.On("Save", input).Return(expected)

	svc := service.NewPostService(mockRepo)

	result := svc.Create(input)

	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := new(mockPostRepository)
	input := model.Post{Title: "Updated"}
	expected := &model.Post{ID: "1", Title: "Updated", Content: "Old", Author: "Author"}

	mockRepo.On("Update", "1", input).Return(expected, true)

	svc := service.NewPostService(mockRepo)

	result, ok := svc.Update("1", input)

	assert.True(t, ok)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdate_NotFound(t *testing.T) {
	mockRepo := new(mockPostRepository)
	input := model.Post{Title: "Updated"}

	mockRepo.On("Update", "404", input).Return(nil, false)

	svc := service.NewPostService(mockRepo)

	result, ok := svc.Update("404", input)

	assert.False(t, ok)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := new(mockPostRepository)
	mockRepo.On("Delete", "1").Return(true)

	svc := service.NewPostService(mockRepo)

	ok := svc.Delete("1")

	assert.True(t, ok)
	mockRepo.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	mockRepo := new(mockPostRepository)
	mockRepo.On("Delete", "999").Return(false)

	svc := service.NewPostService(mockRepo)

	ok := svc.Delete("999")

	assert.False(t, ok)
	mockRepo.AssertExpectations(t)
}

func TestList_Success(t *testing.T) {
	mockRepo := new(mockPostRepository)
	expected := []model.Post{
		{ID: "1", Title: "A"}, {ID: "2", Title: "B"},
	}
	mockRepo.On("FindAll").Return(expected)

	svc := service.NewPostService(mockRepo)

	result := svc.List()

	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}
