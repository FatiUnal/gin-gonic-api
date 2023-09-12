package controller

import (
	"github.com/fatiunal/gin-gonic-copy/entity"
	"github.com/fatiunal/gin-gonic-copy/service"
	"github.com/gin-gonic/gin"
)

type TruckController interface {
	FindAll() []entity.Truck
	Save(ctx *gin.Context) entity.Truck
}

type controller struct {
	truck service.TruckService
}

func New(services service.TruckService) TruckController {
	return &controller{
		truck: services,
	}
}

func (c *controller) FindAll() []entity.Truck {
	return c.truck.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Truck {
	var truck entity.Truck
	ctx.BindJSON(&truck)
	c.truck.Save(truck)
	return truck
}
