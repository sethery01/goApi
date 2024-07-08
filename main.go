package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "HEALTH OKAY",
	})
}

func main() {

	data := Data{
		Number: 20,
	}
	fmt.Printf("%d\n", data.Number)

	router := gin.Default()
	router.GET("/health", healthCheck)
	router.Run() // listen and serve on 0.0.0.0:8080
}
