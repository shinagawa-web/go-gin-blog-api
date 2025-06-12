package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/handler"
	"go-gin-blog-api/service"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	postService := service.NewPostService()
	postHandler := handler.NewPostHandler(postService)
	postHandler.RegisterRoutes(r)

	r.Run(":8080") // ポート8080で起動
}
