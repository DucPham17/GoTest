package controller

import (
	"gotest/src/entity"
	"gotest/src/service"
	"gotest/src/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	videoService service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) Controller {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidatorField)
	return &controller{
		videoService: service,
	}
}

func (con *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	er := ctx.ShouldBindJSON(&video)
	if er != nil {
		return er
	}
	er = validate.Struct(video)
	if er != nil {
		return er
	}
	con.videoService.Save(video)
	return nil
}

func (con *controller) FindAll() []entity.Video {
	return con.videoService.FindAll()
}
