package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/8180149/flutter-artifacts-api/config"

	"github.com/gin-gonic/gin"
)

func DownloadArtifact(c *gin.Context) {
	artifact := c.Param("artifact")
	version := c.Param("version")
	artifactPath := filepath.Join(config.ArtifactDir, artifact, version)

	files, err := ioutil.ReadDir(artifactPath)
	if err != nil || len(files) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No files found for this version"})
		return
	}

	filePath := filepath.Join(artifactPath, files[0].Name())
	c.File(filePath)
}
