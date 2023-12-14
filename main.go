package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoisie/mustache"
	"github.com/minpeter/rctf-backend/api"
)

// ClientConfig는 client-config.json의 구조를 정의합니다.
type ClientConfig struct {
	Meta            Meta              `json:"meta"`
	HomeContent     string            `json:"homeContent"`
	Sponsors        []interface{}     `json:"sponsors"`
	GlobalSiteTag   string            `json:"globalSiteTag"`
	CtfName         string            `json:"ctfName"`
	Divisions       map[string]string `json:"divisions"`
	DefaultDivision string            `json:"defaultDivision"`
	Origin          string            `json:"origin"`
	StartTime       int64             `json:"startTime"`
	EndTime         int64             `json:"endTime"`
	EmailEnabled    bool              `json:"emailEnabled"`
	UserMembers     bool              `json:"userMembers"`
	FaviconUrl      string            `json:"faviconUrl"`
}

// Meta는 client-config.json의 "meta" 부분을 정의합니다.
type Meta struct {
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
}

var clientConfig ClientConfig

func serveIndex(c *gin.Context) {

	rendered := struct {
		JSONConfig string
		Config     ClientConfig
	}{
		JSONConfig: toJSON(clientConfig),
		Config:     clientConfig,
	}

	// Use mustache to render the index.html template
	html := mustache.RenderFile("build/index.html", rendered)
	c.Writer.WriteString(html)
}

func serveFile(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.Writer.WriteHeader(404)
		return
	}

	c.File("build/" + path)
}

func main() {
	loadClientConfig()

	app := api.NewRouter()

	app.GET("/", serveIndex)

	app.GET("/:path", serveFile)

	if err := app.Run("localhost:3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadClientConfig() {
	configFile, err := os.Open("client-config.json")
	if err != nil {
		fmt.Printf("Error opening client-config.json: %v\n", err)
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&clientConfig)
	if err != nil {
		fmt.Printf("Error decoding client-config.json: %v\n", err)
		return
	}
}

func toJSON(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(jsonData)
}