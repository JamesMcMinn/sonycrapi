package main

import (
	"fmt"
	"time"

	"github.com/jamesmcminn/sonycrapi"
)

func main() {
	camera := sonycrapi.NewCamera("http://192.168.122.1:8080/sony/")

	for {
		urls, err := camera.TakePicture()
		if err != nil {
			time.Sleep(time.Second / 4)
		}
		fmt.Println(urls)
	}
}
