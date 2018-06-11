package main

import (
	"log"
	"net/http"

	"github.com/jamesmcminn/sonycrapi"
)

var camera = sonycrapi.NewCamera("http://192.168.122.1:8080/sony/")
var frames = make(chan *sonycrapi.LiveviewPayload, 1)

func main() {
	camera.StopLiveview()
	lv, err := camera.StartLiveviewWithSize(sonycrapi.LiveviewSizes.M)
	if err != nil {
		log.Println(err)
		return
	}

	go lv.Decode(frames)

	go fillFrameBuffers()
	http.HandleFunc("/liveview", liveviewHTTPHandler)
	http.ListenAndServe(":3000", nil)
}
