package main

import (
	"net/http"
	"quickdrop_backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "QuickDrop backend is live!",
		})
	})
	router.GET("/api/manifest", handlers.GetManifest)

	router.POST("/api/driver/location", handlers.UpdateLocation)
	router.Run(":8000")
}
