package main

import (
	"golang/rest-api-gin/controller"
	"golang/rest-api-gin/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

type album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

func main() {
	router := gin.Default()

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(201, VideoController.Save(ctx))
	})

	router.Run("localhost:8080")
}
