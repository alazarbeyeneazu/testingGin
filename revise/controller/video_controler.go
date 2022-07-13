package controller

import (
	"testingme/entity"
	"testingme/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type Adapter struct {
	service service.VideoService
}

func NewAdapter(srvice service.VideoService) *Adapter {
	return &Adapter{service: srvice}
}

func (api Adapter) FindAll() []entity.Video {
	return api.service.FindAll()
}
func (api *Adapter) Save(ctx *gin.Context) entity.Video {
	var Video entity.Video
	ctx.BindJSON(&Video)
	api.service.Save(Video)
	return Video
}
