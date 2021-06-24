package service

import "github.com/DucPham17/GoTest/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(temp entity.Video) entity.Video {
	service.videos = append(service.videos, temp)
	return temp
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
