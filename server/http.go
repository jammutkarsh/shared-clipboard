package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getClips(context *gin.Context) {
	clipData := readJson()
	context.JSON(http.StatusOK, clipData)
}

func postClip(context *gin.Context) {
	var newClip Clipboard
	if err := context.BindJSON(&newClip); err != nil {
		return
	}
	clips := readJson()
	clips = append(clips, newClip)
	writeJson(clips)
	context.JSON(http.StatusCreated, newClip)
}

func deleteClips(context *gin.Context) {
	clear()
	context.JSON(http.StatusOK, readJson())
}

// TODO In a web server, if you panic you miss the opportunity to respond appropriately to the client with say, an http 500
