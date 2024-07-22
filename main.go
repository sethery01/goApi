package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "HEALTH OKAY",
	})
}

func getAlerts(c *gin.Context) {
	url := "https://api.weather.gov/alerts/active/area/MO"

	// Make the request to get weather alerts
	response, error := http.Get(url)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	// Read body of response
	body, error := io.ReadAll(response.Body)
	if error == nil {
		c.JSON(http.StatusOK, gin.H{"data": string(body)})
	}
}

func main() {
	router := gin.Default()
	router.GET("/health", healthCheck)
	router.GET("/alerts", getAlerts)
	router.Run() // listen and serve on 0.0.0.0:8080
}
