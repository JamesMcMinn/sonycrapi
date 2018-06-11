package main

import (
	"runtime"

	"github.com/jamesmcminn/sonycrapi"
)

func main() {
	camera := sonycrapi.NewCamera("http://192.168.122.1:8080/sony/")
	camera.Initilize()

	runtime.Goexit()
}
