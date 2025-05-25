package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func LeetCodeHandler(c *gin.Context) {
	username := os.Getenv("LEETCODE_USERNAME")
	if username == "" {
		username = "hrqmonteiro" // fallback
	}
	// Allow override via ?username=
	if queryUser := c.Query("username"); queryUser != "" {
		username = queryUser
	}

	apiURL := fmt.Sprintf("https://leetcode-stats-api.herokuapp.com/%s", username)
	resp, err := http.Get(apiURL)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.String(http.StatusInternalServerError, "LeetCode API error")
		return
	}

	var data struct {
		TotalSolved        int `json:"totalSolved"`
		Ranking            int `json:"ranking"`
		EasySolved         int `json:"easySolved"`
		MediumSolved       int `json:"mediumSolved"`
		HardSolved         int `json:"hardSolved"`
		ContributionPoints int `json:"contributionPoints"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		c.String(http.StatusInternalServerError, "Failed to decode data")
		return
	}

	html := fmt.Sprintf(`
		<span>Total Solved: %d</span>
		<span>Easy: %d</span>
		<span>Medium: %d</span>
		<span>Hard: %d</span>
		<span>Contest Rank: %d</span>
		<span>Contribution Points: %d</span>
	`, data.TotalSolved, data.EasySolved, data.MediumSolved, data.HardSolved, data.Ranking, data.ContributionPoints)

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
