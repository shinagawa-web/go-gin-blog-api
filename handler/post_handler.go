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
	r.GET("/posts/:id", h.GetPostByID)
	r.PATCH("/posts/:id", h.UpdatePost)
	r.DELETE("/posts/:id", h.DeletePost)
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

func (h *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	post, found := h.postService.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var update model.Post
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post, updated := h.postService.Update(id, update)
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	if deleted := h.postService.Delete(id); !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
