package main

import (
	"ginTutorial/controller"
	"ginTutorial/middleware"
	"ginTutorial/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutPut() {
	f, _ := os.Create("activities.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutPut()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	addr := ":8080"

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"welcome": "To Gin",
	// 	})
	// })
	apiGroup := server.Group("/api")
	{
		apiGroup.GET("/vidoes", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiGroup.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				var know = ctx.ClientIP()
				ctx.JSON(http.StatusBadRequest, gin.H{"result": "fuck you", "you are from => ": know, "why": err.Error()})
				return
			}
			ctx.JSON(200, gin.H{"result": "bravo"})
		})
	}

	viewRoute := server.Group("/view")
	{
		viewRoute.GET("/videos", videoController.ShowAll)
	}
	server.Run(addr)

}
