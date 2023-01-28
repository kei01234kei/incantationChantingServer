// Package util contains a collection of functions used by the server package.
package util

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

// SaveSentFileToLocal saves uploaded file to local tmp directory and returns the saved file name to local.
// The name of saved file is hash generated from file name of origin file.
func SaveSentFileToLocal(c *gin.Context) (string, string, error) {
	file, _ := c.FormFile("file")
	fileName := filepath.Base(file.Filename)
	// Rename the uploaded file using generated hash from file name
	hashedFileName := sha256.Sum256([]byte(fileName))

	savedFileName := fmt.Sprintf("%s%s", hex.EncodeToString(hashedFileName[:]), filepath.Ext(fileName))
	err := c.SaveUploadedFile(file, fmt.Sprintf("./tmp/%s", savedFileName))
	if err != nil {
		return fileName, "", fmt.Errorf("c.SaveUploadedFile: %v", err)
	}
	return fileName, savedFileName, nil
}

// UploadFile uploads the designated file to cloud storage.
func UploadFile(object string) error {
	bucket := "incantation-chanting-server"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	f, err := os.Open(fmt.Sprintf("./tmp/%s", object))
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	o := client.Bucket(bucket).Object(object)
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

// GetObjectURLAndURI returns the url and uri designated in function arguments.
func GetObjectURLAndURI(fileName string) (string, string) {
	URL := fmt.Sprintf("https://storage.cloud.google.com/incantation-chanting-server/%s", fileName)
	URI := fmt.Sprintf("gs://incantation-chanting-server/%s", fileName)
	return URL, URI
}
