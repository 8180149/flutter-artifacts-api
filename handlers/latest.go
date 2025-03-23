package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/8180149/flutter-artifacts-api/config"
	"github.com/gin-gonic/gin"
)

func GetLatestVersion(c *gin.Context) {
	artifactDirs, err := os.ReadDir(config.ArtifactDir)
	if err != nil || len(artifactDirs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No artifacts found"})
		return
	}

	sort.Slice(artifactDirs, func(i, j int) bool {
		infoI, _ := artifactDirs[i].Info()
		infoJ, _ := artifactDirs[j].Info()
		return infoI.ModTime().After(infoJ.ModTime())
	})

	latestArtifact := artifactDirs[0].Name()
	latestArtifactPath := filepath.Join(config.ArtifactDir, latestArtifact)
	versionDirs, err := os.ReadDir(latestArtifactPath)
	if err != nil || len(versionDirs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No versions found for latest artifact"})
		return
	}

	sort.Slice(versionDirs, func(i, j int) bool {
		infoI, _ := versionDirs[i].Info()
		infoJ, _ := versionDirs[j].Info()
		return infoI.ModTime().After(infoJ.ModTime())
	})

	latestVersion := versionDirs[0].Name()
	downloadURL := "/download/" + latestArtifact + "/" + latestVersion

	c.JSON(http.StatusOK, gin.H{
		"latest_artifact": latestArtifact,
		"latest_version":  latestVersion,
		"download_url":    downloadURL,
	})
}
