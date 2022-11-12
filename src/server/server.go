// Package server is to return something when api is invoked.
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func GetFileTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.File(fmt.Sprintf("./tmp/%s", c.Param("name")))
	}
}
