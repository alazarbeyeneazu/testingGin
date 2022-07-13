package main

import (
	"testingme/controller"
	"testingme/middleware"
	"testingme/service"

	"github.com/gin-gonic/gin"
)

var (
	vservice    service.VideoService       = service.NewAdapter()
	vcontroller controller.VideoController = controller.NewAdapter(vservice)
)

func main() {
	server := gin.New()
	server.Use(middleware.Logger())
	addr := ":8080"
	server.GET("/vidoes", func(ctx *gin.Context) {
		ctx.JSON(200, vcontroller.FindAll())
	})
	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(200, vcontroller.Save(ctx))

	})
	server.Run(addr)
}
