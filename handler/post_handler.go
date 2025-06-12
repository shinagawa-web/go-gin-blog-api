package handler

import (
	"go-gin-blog-api/model"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// 簡易的なメモリ保存（本来はDBを使う）
var (
	posts = []model.Post{}
	mutex = &sync.Mutex{}
)

func RegisterPostRoutes(r *gin.Engine) {
	r.POST("/posts", CreatePost)
	r.GET("/posts", GetPosts)
}

// 記事の投稿ハンドラ
func CreatePost(c *gin.Context) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutex.Lock()
	posts = append(posts, newPost)
	mutex.Unlock()

	c.JSON(http.StatusCreated, newPost)
}

// 記事一覧取得ハンドラ
func GetPosts(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	c.JSON(http.StatusOK, posts)
}
