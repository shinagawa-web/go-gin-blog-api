package handler_test

import (
	"go-gin-blog-api/handler"
	"go-gin-blog-api/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// サービスをモック化
type mockPostService struct {
	mock.Mock
}

func NewMockPostService() *mockPostService {
	return &mockPostService{}
}

func (m *mockPostService) GetByID(id string) (*model.Post, bool) {
	args := m.Called(id)
	if post := args.Get(0); post != nil {
		return post.(*model.Post), true
	}
	return nil, false
}

func (m *mockPostService) Create(post model.Post) *model.Post {
	args := m.Called(post)
	if p := args.Get(0); p != nil {
		return p.(*model.Post)
	}
	return nil
}

func (m *mockPostService) Update(id string, post model.Post) (*model.Post, bool) {
	args := m.Called(id, post)
	if p := args.Get(0); p != nil {
		return p.(*model.Post), true
	}
	return nil, false
}

func (m *mockPostService) Delete(id string) bool {
	args := m.Called(id)
	return false != args.Bool(0)
}

func (m *mockPostService) List() []model.Post {
	args := m.Called()
	if list := args.Get(0); list != nil {
		return list.([]model.Post)
	}
	return nil
}

func TestGetPost_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := NewMockPostService()
	mockSvc.On("GetByID", "1").Return(&model.Post{
		ID: "1", Title: "Hello", Content: "Test", Author: "Alice",
	}, nil)

	h := handler.NewPostHandler(mockSvc)

	r := gin.Default()
	r.GET("/posts/:id", h.GetPostByID)

	req, _ := http.NewRequest("GET", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id":"1","title":"Hello","content":"Test","author":"Alice"}`, w.Body.String())

	mockSvc.AssertExpectations(t)
}

func TestGetPost_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := NewMockPostService()
	mockSvc.On("GetByID", "999").Return(nil, false)

	h := handler.NewPostHandler(mockSvc)

	r := gin.Default()
	r.GET("/posts/:id", h.GetPostByID)

	req, _ := http.NewRequest("GET", "/posts/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"error":"Post not found"}`, w.Body.String())
}
