package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// DB Json File Check
	if !FileExists(fileLocation) {
		_, err := os.Create(fileLocation)
		clear()
		if err != nil {
			log.Fatal(err)
		}
	}
	// Init router
	router := gin.Default()

	router.GET("/fetch", getClips)
	router.POST("/add", postClip)
	router.DELETE("/clear", deleteClips)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
