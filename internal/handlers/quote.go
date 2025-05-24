package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Quote struct {
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

func QuoteHandler(c *gin.Context) {
	cwd, _ := os.Getwd()
	quotePath := filepath.Join(cwd, "content", "quotes", "quotes.json")

	file, err := os.Open(quotePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open quotes file")
		return
	}
	defer file.Close()

	var quotes Quotes
	if err := json.NewDecoder(file).Decode(&quotes); err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode quotes file")
		return
	}

	if len(quotes.Quotes) == 0 {
		c.String(http.StatusInternalServerError, "No quotes found")
		return
	}

	rand.Seed(time.Now().UnixNano())
	quote := quotes.Quotes[rand.Intn(len(quotes.Quotes))]

	response := quote.Quote + " – " + quote.Character
	c.String(http.StatusOK, response)
}
