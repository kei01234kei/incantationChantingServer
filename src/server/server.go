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

func UploadFileTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, fmt.Sprintf("./tmp/%s", file.Filename))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}
}

func GetFileTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.File(fmt.Sprintf("./tmp/%s", c.Param("name")))
	}
}
