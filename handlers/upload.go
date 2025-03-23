package handlers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/8180149/flutter-artifacts-api/config"
	"github.com/gin-gonic/gin"
)

func UploadArtifact(c *gin.Context) {
	osSystem := c.PostForm("os") // Receiving OS instead of artifact
	version := c.PostForm("version")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
		return
	}

	// Define the path as artifacts/{os}/{version}.zip
	filePath := filepath.Join(config.ArtifactDir, osSystem, version+".zip")

	// Ensure the parent directory exists
	dirPath := filepath.Join(config.ArtifactDir, osSystem)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directories"})
		return
	}

	// Save the uploaded file
	if err := saveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "path": filePath})
}

// Helper function to save uploaded file
func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, src)
	return err
}
