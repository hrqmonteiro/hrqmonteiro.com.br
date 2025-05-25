package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GitHubHandler(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		username = "hrqmonteiro"
	}

	// Fetch user data
	userURL := fmt.Sprintf("https://api.github.com/users/%s", username)
	userResp, err := http.Get(userURL)
	if err != nil || userResp.StatusCode != http.StatusOK {
		c.String(http.StatusInternalServerError, "Error fetching user data")
		return
	}
	defer userResp.Body.Close()

	var user struct {
		Followers int       `json:"followers"`
		CreatedAt time.Time `json:"created_at"`
	}
	if err := json.NewDecoder(userResp.Body).Decode(&user); err != nil {
		c.String(http.StatusInternalServerError, "Error decoding user data")
		return
	}

	// Fetch repositories
	reposURL := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username)
	reposResp, err := http.Get(reposURL)
	if err != nil || reposResp.StatusCode != http.StatusOK {
		c.String(http.StatusInternalServerError, "Error fetching repositories")
		return
	}
	defer reposResp.Body.Close()

	var repositories []struct {
		StargazersCount int `json:"stargazers_count"`
	}
	if err := json.NewDecoder(reposResp.Body).Decode(&repositories); err != nil {
		c.String(http.StatusInternalServerError, "Error decoding repositories")
		return
	}

	repoCount := len(repositories)
	starCount := 0
	for _, repo := range repositories {
		starCount += repo.StargazersCount
	}

	createdAtFormatted := user.CreatedAt.Format("January 2, 2006") // e.g., "May 25, 2025"

	// Build HTML response
	html := strings.Join([]string{
		fmt.Sprintf("<span>%d followers</span>", user.Followers),
		fmt.Sprintf("<span>%d repositories</span>", repoCount),
		fmt.Sprintf("<span>%d stars</span>", starCount),
		fmt.Sprintf("<span>Joined on %s</span>", createdAtFormatted),
	}, "\n")

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, html)
}
