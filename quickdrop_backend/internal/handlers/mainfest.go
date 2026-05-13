package handlers

import (
	"net/http"
	"quickdrop_backend/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetManifest(c *gin.Context) {
	dummyPackages := []models.Package{
		{
			ID:        primitive.NewObjectID(),
			Address:   "A-402, Pearl Heights",
			Latitude:  28.4595,
			Longitude: 77.0266,
			Status:    "pending",
		},
		{
			ID:        primitive.NewObjectID(),
			Address:   "Shop 4, City Center Mall",
			Latitude:  28.4721,
			Longitude: 77.0453,
			Status:    "pending",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dummyPackages,
	})
}
