package server

import (
	"fmt"
	"go-gin-blog-api/config"
	"go-gin-blog-api/handler"
	"go-gin-blog-api/logger"
	"go-gin-blog-api/repository"
	"go-gin-blog-api/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New() (*gin.Engine, error) {
	if err := logger.Init(config.AppConfig.Env); err != nil {
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(GinZapMiddleware())

	postRepo := repository.NewPostRepository()
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)
	postHandler.RegisterRoutes(r)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r, nil
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
