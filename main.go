package main

import (
	"fmt"
	"go-gin-blog-api/config"
	"go-gin-blog-api/handler"
	"go-gin-blog-api/logger"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	if err := logger.Init(cfg.Env); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Log.Sync()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(GinZapMiddleware())

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

func GinZapMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		logger.Log.Info("request completed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
		)
	}
}
