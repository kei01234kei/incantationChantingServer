package main

import (
	"github.com/gin-gonic/gin"
	"incantationChantingServer/src/server"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/test", server.GetTest)
	router.GET("/test-get-file/:name", server.GetFileTest)
	router.GET("/convert", server.ConvertSoundToText)
	router.POST("/test-upload-file", server.UploadFileTest)
	router.POST("/upload", server.UploadFile)
	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
