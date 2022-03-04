package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// DB Json File Check
	if !FileExists(fileLocation) {
		_, err := os.Create(fileLocation)
		clear()
		if err != nil {
			panic(err)
		}
	}
	// Init router
	router := gin.Default()

	router.GET("/fetch", getClips)
	router.POST("/add", postClip)
	router.DELETE("/clear", deleteClips)

	err := router.Run("localhost:8080")
	errHanding(err)
}
