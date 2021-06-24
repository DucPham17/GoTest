package main

import (
	"github.com/DucPham17/GoTest/controller"
	"github.com/DucPham17/GoTest/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()

	control controller.Controller = controller.New(videoService)
)

func main() {
	r := gin.Default()
	r.GET("/videos", func(c *gin.Context) {
		c.JSON(200, control.FindAll())
	})
	r.POST("/videos", func(c *gin.Context) {
		c.JSON(200, control.Save(c))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
