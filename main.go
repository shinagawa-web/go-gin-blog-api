package main

import (
	"fmt"
	"go-gin-blog-api/config"
	"go-gin-blog-api/internal/server"
	"go-gin-blog-api/logger"
	"log"
)

func main() {
	config.Load()
	r, err := server.New()
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Log.Sync()

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	log.Printf("Server running on %s (%s)", addr, config.AppConfig.Env)
	r.Run(addr)
}
