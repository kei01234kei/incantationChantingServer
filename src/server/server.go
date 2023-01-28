// Package server is to return something when api is invoked.
package server

import (
	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"incantationChantingServer/src/util"
	"log"
	"net/http"
	"path/filepath"
)

func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func GetFileTest(c *gin.Context) {
	c.File(fmt.Sprintf("./tmp/%s", c.Param("name")))
}

func UploadFileTest(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	err := c.SaveUploadedFile(file, fmt.Sprintf("./tmp/%s", file.Filename))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to upload '%s' !", file.Filename))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("[Success]: '%s' uploaded!", file.Filename))
	}
}

func UploadFile(c *gin.Context) {
	fileName, savedFileName, err := util.SaveSentFileToLocal(c)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to save hashed file '%s' to server ! %v", savedFileName, err))
		return
	}
	err = util.UploadFile(savedFileName)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to uploading file '%s' to cloud storage ! %v", savedFileName, err))
		return
	}
	fileURL, fileURI := util.GetObjectURLAndURI(savedFileName)
	c.JSON(http.StatusOK, gin.H{
		"name":            fileName,
		"saved_file_name": savedFileName,
		"url":             fileURL,
		"uri":             fileURI,
	})
}

func ConvertSoundToText(c *gin.Context) {
	bucket := "incantation-chanting-server"
	fileName := c.Query("filename")
	if fileName == "" {
		c.String(http.StatusInternalServerError, "[Error]: Send filename using filename query !")
		return
	}
	if filepath.Ext(fileName) != ".wav" {
		c.String(http.StatusInternalServerError, "[Error]: Enter wav file in file name query !")
		return
	}
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to create client: %v", err))
		return
	}
	defer client.Close()
	fileURI := fmt.Sprintf("gs://%s/%s", bucket, fileName)
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:     speechpb.RecognitionConfig_LINEAR16,
			LanguageCode: "ja-JP",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: fileURI},
		},
	})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("[Error]: Failed to recognize: %v", err))
		return
	}
	if len(resp.Results) == 0 {
		c.String(http.StatusBadRequest, "[Error]: Failed to recognize sound !")
		return
	}
	c.JSON(http.StatusOK, resp)
}
