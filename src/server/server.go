// Package server is to return something when api is invoked.
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"incantationChantingServer/src/util"
	"log"
	"net/http"
)

func GetTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	}
}

func GetFileTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.File(fmt.Sprintf("./tmp/%s", c.Param("name")))
	}
}

func UploadFileTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, fmt.Sprintf("./tmp/%s", file.Filename))
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to upload '%s' !", file.Filename))
		} else {
			c.String(http.StatusOK, fmt.Sprintf("[Success]: '%s' uploaded!", file.Filename))
		}
	}
}

func UploadFile() func(c *gin.Context) {
	return func(c *gin.Context) {
		fileName, hashedFileName, err := util.SaveSentFileToLocal(c)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to save hashed file '%s' to server ! %v", hashedFileName, err))
			return
		}
		err = util.UploadFile(hashedFileName)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to uploading file '%s' to cloud storage ! %v", hashedFileName, err))
			return
		}
		fileURL, fileURI := util.GetObjectURLAndURI(hashedFileName)
		c.JSON(http.StatusOK, gin.H{
			"name":        fileName,
			"hashed_name": hashedFileName,
			"url":         fileURL,
			"uri":         fileURI,
		})
	}
}
