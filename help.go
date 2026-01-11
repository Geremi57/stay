package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var CLIENT_ID = "tFsns0K9BMeuQ0xwN9iy1nWno6tTjGWxkeFOmLx9ZIw"
var CLIENT_SECRET = "efKNWyiX9G6kLVbZEc6Rt8Y42hQ7MjK4smhFvhqud94"

func GetToken() (string, error) {
	url := "https://api.throughlinecare.com/oauth/token"

	payload := []byte(`{
		"grant_type": "client_credentials",
		"client_id": "` + CLIENT_ID + `",
		"client_secret": "` + CLIENT_SECRET + `"
	}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	token := result["access_token"].(string)
	return token, nil
}

// ==================== GET HELPLINES ====================
func GetHelplines(c *gin.Context) {
	country := c.Query("country")

	if country == "" {
		country = "US" // default
	}

	token, err := GetToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	reqUrl := fmt.Sprintf(
		"https://api.throughlinecare.com/v1/helplines?country_code=%s&limit=20",
		country,
	)

	req, _ := http.NewRequest("GET", reqUrl, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request error"})
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	c.Data(http.StatusOK, "application/json", body)
}

// ==================== SERVER ====================
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api/helplines", GetHelplines)
	r.Run(":9090")
}
