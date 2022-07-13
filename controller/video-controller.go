package controller

import (
	"ginTutorial/entity"
	"ginTutorial/service"
	"net/http"

	customValidators "ginTutorial/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) *controller {
	validate = validator.New()
	validate.RegisterValidation("is-cool", customValidators.ValidateCoolTitles)
	return &controller{service: service}
}
func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBind(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
func (controller *controller) ShowAll(ctx *gin.Context) {
	video := controller.service.FindAll()
	data := gin.H{
		"title": "Videos page ",
		"video": video,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
