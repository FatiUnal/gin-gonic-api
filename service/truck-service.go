package service

import "github.com/fatiunal/gin-gonic-copy/entity"

type TruckService interface {
	FindAll() []entity.Truck
	Save(entity.Truck) entity.Truck
}

type truck struct {
	trucks []entity.Truck
}

func New() TruckService {
	return &truck{}
}

func Asd() string {
	return "asd"
}

func (truckService *truck) FindAll() []entity.Truck {
	return truckService.trucks
}

func (truckService *truck) Save(truck entity.Truck) entity.Truck {
	truckService.trucks = append(truckService.trucks, truck)
	return truck
}
