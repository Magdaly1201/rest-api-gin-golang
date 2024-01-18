package main

import (
	"golang/rest-api-gin/controller"
	"golang/rest-api-gin/middleware"
	"golang/rest-api-gin/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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
		ctx.JSON(201, VideoController.Save(ctx))
	})

	router.Run("localhost:8080")
}
