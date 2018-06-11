package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/jamesmcminn/sonycrapi"
)

var camera = sonycrapi.NewCamera("http://192.168.122.1:8080/sony/")

func main() {

	log.Println("Initilizing camera...")
	err := camera.Initilize()
	if err != nil {
		log.Println(err)
	}

	log.Println("Setting shoot mode...")
	err = camera.SetShootMode(sonycrapi.ShootModes.Still)
	if err != nil {
		log.Println(err)
	}

	cameraSetup()

	log.Println("Starting live view...")
	go startLiveView()

	http.HandleFunc("/liveview", liveviewHTTPHandler)
	http.HandleFunc("/takephoto", takephotoHTTPHandler)
	http.Handle("/print/", printHandler{})
	http.Handle("/photo/", photoHandler{})
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)

	runtime.Goexit()
}

func startLiveView() {
	camera.StopLiveview()
	lv, err := camera.StartLiveviewWithSize(sonycrapi.LiveviewSizes.L)
	if err != nil {
		log.Println(err)
		return
	}

	go lv.Decode(frames)
	go fillFrameBuffers()

}

func cameraSetup() {
	err := camera.SetContinuousShootingMode(sonycrapi.ContinuousShootingModes.Single)
	if err != nil {
		fmt.Println(err)
	}

	err = camera.SetShootMode(sonycrapi.ShootModes.Still)
	if err != nil {
		log.Println("Shoot mode err:", err)
	}

	err = camera.SetStillQuality(sonycrapi.StillQualities.RAWJPG)
	if err != nil {
		log.Println(err)
	}

	err = camera.SetPostviewImageSize(sonycrapi.PostViewSizes.Original)
	if err != nil {
		log.Println(err)
	}

	camera.UpdateState(false)
}
