package main

import (
	"fmt"
	"go-gin-blog-api/config"
	"go-gin-blog-api/handler"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	postRepo := repository.NewPostRepository()
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)
	postHandler.RegisterRoutes(r)

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	log.Printf("Server running on %s (%s)", addr, config.AppConfig.Env)
	r.Run(addr)
}
