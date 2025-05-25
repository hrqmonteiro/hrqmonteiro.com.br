package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type SpotifyTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type SpotifyArtist struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
}

type SpotifyImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type SpotifyAlbum struct {
	Name   string         `json:"name"`
	Images []SpotifyImage `json:"images"`
}

type SpotifyTrack struct {
	Album        SpotifyAlbum      `json:"album"`
	Artists      []SpotifyArtist   `json:"artists"`
	ExternalUrls map[string]string `json:"external_urls"`
	Name         string            `json:"name"`
}

type SpotifyNowPlaying struct {
	Item      SpotifyTrack `json:"item"`
	IsPlaying bool         `json:"is_playing"`
}

type SpotifyResponse struct {
	Album         string `json:"album"`
	AlbumImageURL string `json:"albumImageUrl"`
	Artist        string `json:"artist"`
	IsPlaying     bool   `json:"isPlaying"`
	SongURL       string `json:"songUrl"`
	Title         string `json:"title"`
	URL           string `json:"url"`
}

const (
	NOW_PLAYING_ENDPOINT = "https://api.spotify.com/v1/me/player/currently-playing"
	TOKEN_ENDPOINT       = "https://accounts.spotify.com/api/token"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}
}

func getAccessToken() (*SpotifyTokenResponse, error) {
	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || refreshToken == "" {
		return nil, fmt.Errorf("missing Spotify environment variables")
	}

	// Create basic auth header
	auth := clientID + ":" + clientSecret
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	// Prepare form data
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", TOKEN_ENDPOINT, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+basicAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp SpotifyTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func getNowPlaying(accessToken string) (*http.Response, error) {
	req, err := http.NewRequest("GET", NOW_PLAYING_ENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	return client.Do(req)
}

func SpotifyHandler(c *gin.Context) {
	// Get access token
	tokenResp, err := getAccessToken()
	if err != nil {
		c.String(http.StatusOK, "Not playing")
		return
	}

	// Get now playing
	resp, err := getNowPlaying(tokenResp.AccessToken)
	if err != nil {
		c.String(http.StatusOK, "Not playing")
		return
	}
	defer resp.Body.Close()

	// Handle no content or error status
	if resp.StatusCode == 204 || resp.StatusCode >= 400 {
		c.String(http.StatusOK, "Not playing")
		return
	}

	var nowPlaying SpotifyNowPlaying
	if err := json.NewDecoder(resp.Body).Decode(&nowPlaying); err != nil {
		c.String(http.StatusOK, "Not playing")
		return
	}

	// Extract artist names
	var artistNames []string
	for _, artist := range nowPlaying.Item.Artists {
		artistNames = append(artistNames, artist.Name)
	}
	artistString := strings.Join(artistNames, ", ")

	// Return formatted string
	fullTitle := fmt.Sprintf("%s - %s", artistString, nowPlaying.Item.Name)
	if len([]rune(fullTitle)) > 50 {
		runes := []rune(fullTitle)
		fullTitle = string(runes[:47]) + "..."
	}
	c.String(http.StatusOK, fullTitle)
}
