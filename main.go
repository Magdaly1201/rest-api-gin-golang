package main

import (
	"golang/rest-api-gin/controller"
	"golang/rest-api-gin/middleware"
	"golang/rest-api-gin/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	router := gin.New()

	router.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	router.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "video input is valid",
			})
		}
	})

	router.Run("localhost:8080")
}
