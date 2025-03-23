package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"

	"github.com/8180149/flutter-artifacts-api/config"
	"github.com/gin-gonic/gin"
)

func GetLatestVersion(c *gin.Context) {
	artifact := c.Param("artifact")
	artifactPath := filepath.Join(config.ArtifactDir, artifact)

	dirs, err := ioutil.ReadDir(artifactPath)
	if err != nil || len(dirs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No versions found"})
		return
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].ModTime().After(dirs[j].ModTime())
	})

	latestVersion := dirs[0].Name()
	downloadURL := "/download/" + artifact + "/" + latestVersion

	c.JSON(http.StatusOK, gin.H{"latest_version": latestVersion, "download_url": downloadURL})
}
