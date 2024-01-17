package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "OK!!",
		})
	})

	router.Run("localhost:8080")
}
