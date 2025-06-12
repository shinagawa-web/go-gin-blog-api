package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-blog-api/handler"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	handler.RegisterPostRoutes(r)

	r.Run(":8080") // ポート8080で起動
}
