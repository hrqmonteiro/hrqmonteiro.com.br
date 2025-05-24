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
	r.GET("/quote", handlers.QuoteHandler)

	fmt.Println("Server starting on :8080")
	r.Run(":8080")
}
