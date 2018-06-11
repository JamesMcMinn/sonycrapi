package sonycrapi

import (
	"log"

	"github.com/huin/goupnp"
)

const upnpSearchString = "urn:schemas-sony-com:service:ScalarWebAPI:1"
const upnpAll = "ssdp:all"

func DiscoverDevice() (camera *Camera, err error) {
	devices, err := goupnp.DiscoverDevices(upnpSearchString)
	if err != nil {
		log.Println(err)
	}
	log.Println("Devices2:", devices)

	if len(devices) > 0 {
		d := devices[0]
		log.Printf("%+v", d.Root)
	}

	return
}
