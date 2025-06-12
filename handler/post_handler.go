package handler

import (
	"go-gin-blog-api/model"
	"go-gin-blog-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{postService: s}
}

func (h *PostHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/posts", h.CreatePost)
	r.GET("/posts", h.GetPosts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := h.postService.Create(newPost)
	c.JSON(http.StatusCreated, created)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts := h.postService.List()
	c.JSON(http.StatusOK, posts)
}
