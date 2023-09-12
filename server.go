package main

import (
	"github.com/fatiunal/gin-gonic-copy/controller"
	"github.com/fatiunal/gin-gonic-copy/service"
	"github.com/gin-gonic/gin"
)

var (
	truckService    service.TruckService       = service.New()
	truckController controller.TruckController = controller.New(truckService)
)

func main() {
	server := gin.Default()

	server.GET("/truck", func(ctx *gin.Context) {
		ctx.JSON(200, truckController.FindAll())
	})

	server.POST("/truck", func(ctx *gin.Context) {
		ctx.JSON(200, truckController.Save(ctx))
	})

	server.Run(":8080")
}
