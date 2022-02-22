package main

import (
	"github.com/gin-gonic/gin"
)

func getReq(ctx *gin.Context) {
	jsonData := readJson()
	ctx.JSON(200, gin.H{
		jsonData,
	})
	err := ctx.BindJSON(jsonData)
	if err != nil {
		panic(err)
	}
}
