package controller

import (
	"github.com/DucPham17/GoTest/entity"
	"github.com/DucPham17/GoTest/service"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Save() entity.Video
	FindAll() []entity.Video
}

type controller struct {
	videoService service.VideoService
}

func New(service service.VideoService) Controller {
	return &controller{
		videoService: service,
	}
}

func (con *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	con.videoService.Save(video)
	return video
}

func (con *controller) FindAll() []entity.Video {
	return con.videoService.FindAll()
}
