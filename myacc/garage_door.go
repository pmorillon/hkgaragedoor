package myacc

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//GarageDoor struct
type GarageDoor struct {
	*accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

//NewGarageDoor return a garage door accessory
func NewGarageDoor(info accessory.Info) *GarageDoor {
	acc := GarageDoor{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddService(acc.GarageDoorOpener.Service)

	return &acc
}
