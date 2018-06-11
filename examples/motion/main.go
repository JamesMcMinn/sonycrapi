package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"

	"github.com/jamesmcminn/sonycrapi"
	"github.com/tajtiattila/blur"

	_ "net/http/pprof"

	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"time"
)

type FramePair struct {
	Original image.Image
	Current  image.Image
	Previous image.Image
	Output   image.Image
	InTime   int64
}

var frames = make(chan *sonycrapi.LiveviewPayload, 1)
var diffFrames = make(chan image.Image)
var absQueue = make(chan *FramePair, 2)
var blurQueue = make(chan *FramePair, 2)
var thresholdQueue = make(chan *FramePair, 2)
var hotQueue = make(chan *FramePair, 2)

var takePhotoQueue = make(chan int64, 1000000)

var camera = sonycrapi.NewCamera("http://192.168.122.1:8080/sony/")

func main() {
	camera.StopLiveview()
	lv, err := camera.StartLiveviewWithSize(sonycrapi.LiveviewSizes.M)
	if err != nil {
		log.Println(err)
		return
	}

	go lv.Decode(frames)

	go takePhotos()

	go runAbsDiff(absQueue, blurQueue)
	go runAbsDiff(absQueue, blurQueue)
	go runBlur(blurQueue, thresholdQueue)
	go runBlur(blurQueue, thresholdQueue)
	go runThreshold(thresholdQueue, hotQueue)
	go runThreshold(thresholdQueue, hotQueue)

	go runHotCheck(hotQueue)

	go runServer()

	var previousFrame *sonycrapi.LiveviewPayload
	var previousImage image.Image
	for {
		frame := <-frames
		buff := bytes.NewBuffer(frame.JPEGData)
		img, _, err := image.Decode(buff)
		var current image.Image
		if err != nil {
			log.Println(err)
			continue
		}

		if previousFrame != nil {
			start := time.Now().UnixNano()
			current = blur.Gaussian(img, 9, blur.ReuseSrc)

			pair := &FramePair{
				Original: img,
				Current:  current,
				Previous: previousImage,
				InTime:   start,
			}

			if len(absQueue) == cap(absQueue) {
				<-absQueue
			}
			absQueue <- pair
		}

		if previousFrame == nil {
			previousImage = blur.Gaussian(img, 9, blur.ReuseSrc)
		} else {
			previousImage = current
		}
		previousFrame = frame
	}
}

func takePhotos() {
	takePhotosUntil := time.Now().Unix() - 1
	for {

		if takePhotosUntil < time.Now().Unix() {
			newtime := <-takePhotoQueue
			if newtime > takePhotosUntil {
				takePhotosUntil = newtime
			}
			continue
		} else {
			for i := 0; i < len(takePhotoQueue); i++ {
				newtime := <-takePhotoQueue
				if newtime > takePhotosUntil {
					takePhotosUntil = newtime
				}
			}
		}

		log.Printf("Taking Photo. %d seconds remaining...\n", takePhotosUntil-time.Now().Unix())

		urls, err := camera.TakePicture()
		if err != nil {
			log.Println(err)
		}

		log.Println(urls)
	}

}

func runServer() {
	go fillFrameBuffers()
	http.HandleFunc("/liveview", liveviewHTTPHandler)
	http.ListenAndServe(":3000", nil)
}

func runHotCheck(in chan *FramePair) {
	for {
		pair := <-in

		hot, pct := HotPixels(pair.Output)

		end := time.Now().UnixNano()
		ttlTime := (end - pair.InTime) / 1e6

		if ttlTime > 500 {
			log.Println(hot, pct, ttlTime, "ms")
		}

		if hot > 225 {
			takePhotoQueue <- (int64(time.Now().Unix()) + 5)

			log.Printf("detected a distance %d / %3.2f%%", hot, pct)

			dest, err := os.Create(fmt.Sprintf("/home/james/motion/%d-%d.png", time.Now().UnixNano(), hot))
			if err != nil {
				log.Println(err)
				continue
			}
			png.Encode(dest, pair.Original)

			destDiff, err := os.Create(fmt.Sprintf("/home/james/motion/%d-%d-diff.png", time.Now().UnixNano(), hot))
			if err != nil {
				log.Println(err)
				continue
			}
			png.Encode(destDiff, pair.Output)

			// diffFrames <- Extract(pair.Original, pair.Output)
		}
	}
}

func runAbsDiff(in, out chan *FramePair) {
	for {
		pair := <-in
		diffFrames <- pair.Original
		pair.Output = AbsDiff(pair.Previous, pair.Current)
		// diffFrames <- pair.Output
		out <- pair
	}
}

func runBlur(in, out chan *FramePair) {
	for {
		pair := <-in
		pair.Output = blur.Gaussian(pair.Output, 6, blur.ReuseSrc)
		out <- pair
	}
}

func runThreshold(in, out chan *FramePair) {
	for {
		pair := <-in
		pair.Output = Threshold(pair.Output, 10)
		out <- pair
	}
}
