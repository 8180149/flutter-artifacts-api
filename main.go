package main

import (
	"github.com/8180149/flutter-artifacts-api/config"
	"github.com/8180149/flutter-artifacts-api/handlers"
	"github.com/8180149/flutter-artifacts-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.POST("/upload", handlers.UploadArtifact)
	r.GET("/latest", handlers.GetLatestVersion) // Updated to remove artifact param
	r.GET("/download/:artifact/:version", handlers.DownloadArtifact)

	r.Run(":" + config.Port)
}
