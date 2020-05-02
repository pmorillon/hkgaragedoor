package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/characteristic"
	"github.com/pmorillon/hkgaragedoor/myacc"
)

var (
	accessoryPinArg = flag.String("pin", "11122333", "HomeKit accessory pin.")
	debugArg = flag.Bool("debug", false, "Debug mode.")
)

func main() {
	fmt.Println("Starting HomeKit Garage Door controller...")
	flag.Parse()

	if *debugArg {
		log.Debug.Enable()
	}

	info := accessory.Info{
		Name: "Garage Door",
		Model: "Raspberry Pi",
		Manufacturer: "Sorillon",
	}

	acc := myacc.NewGarageDoor(info)

	t, err := hc.NewIPTransport(hc.Config{Pin: *accessoryPinArg}, acc.Accessory)
	if err != nil {
		log.Info.Panic(err)
	}

	acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)

	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(target int) {
		if target == characteristic.TargetDoorStateClosed {
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosing)
			time.Sleep(10 * time.Second)
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
		} else if target == characteristic.TargetDoorStateOpen {
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpening)
			time.Sleep(10 * time.Second)
			acc.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpen)
		}
	})

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}