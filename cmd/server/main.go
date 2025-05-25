package main

import (
	"fmt"
	"hrqmonteiro.com.br/internal/handlers"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Working directory: %s\n", cwd)

	staticPath := filepath.Join(cwd, "static")
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		staticPath = filepath.Join(cwd, "../static")
		if _, err := os.Stat(staticPath); os.IsNotExist(err) {
			staticPath = filepath.Join(cwd, "../../static")
		}
	}

	fmt.Printf("Static path: %s\n", staticPath)

	r.Static("/static", staticPath)
	r.GET("/", handlers.IndexHandler)
	r.GET("/github", handlers.GitHubHandler)
	r.GET("/leetcode", handlers.LeetCodeHandler)
	r.GET("/spotify", handlers.SpotifyHandler)
	r.GET("/quote", handlers.QuoteHandler)

	err = r.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server starting on :8000")
}
