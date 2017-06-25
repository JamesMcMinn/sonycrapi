package sonycrapi

import (
	"fmt"
	"log"

	"github.com/huin/goupnp"
)

const upnpSearchString = "urn:schemas-sony-com:service:ScalarWebAPI:1"

func DiscoverDevice() (camera *Camera, err error) {
	devices, err := goupnp.DiscoverDevices(upnpSearchString)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(devices)

	return
}
