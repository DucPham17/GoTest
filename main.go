package main

import (
	"fmt"
	"gotest/src/controller"
	"gotest/src/entity"
	"gotest/src/middlewares"
	"gotest/src/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()

	control controller.Controller = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("go.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	fmt.Println(entity.GetString())
	r := gin.Default()
	r.Use(middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	route := r.Group("/api")
	route.GET("/videos", func(c *gin.Context) {
		c.JSON(200, control.FindAll())
	})
	route.POST("/videos", func(c *gin.Context) {
		err := control.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"message": "Good!"})
		}
	})
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
