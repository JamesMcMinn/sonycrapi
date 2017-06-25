package sonycrapi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// LiveviewSize defined a pre-defined resolution or quality for the live view stream. Instances
// of LiveviewSize are defined in the LiveviewSizes struct variable.
type LiveviewSize string

// LiveviewSizes is defined as above.
var LiveviewSizes = struct {
	L LiveviewSize
	M LiveviewSize
}{"L", "M"}

var liveviewImagePayloadCode byte = 0x01
var liveviewInfoPayloadCode byte = 0x02
var liveviewCommonHeaderStartCode byte = 0xff
var payloadHeaderStartCode = []byte{0x24, 0x35, 0x68, 0x79}

// Liveview defined a live view stream and it's basic instance properties
type Liveview struct {
	URL          string
	Stream       *bufio.Reader
	HTTPResponse *http.Response
	Camera       *Camera
}

type LiveviewPayload struct {
	LiveviewCommonHeader

	PayloadSize uint32
	PaddingSize uint8

	*LiveviewInfoPayload
	*LiveviewImagePayload
}

type LiveviewCommonHeader struct {
	PayloadType    byte
	SequenceNumber uint16
	Timestamp      uint32
}

type LiveviewImagePayload struct {
	Reserved [120]byte // Who the fuck knows what Sony are using this for. It's not documente
	JPEGData []byte
}

type LiveviewInfoPayload struct {
	VersionMajor    uint8
	VersionMinor    uint8
	FrameCount      uint16
	SingleFrameSize uint16
	Reserved        [114]byte // Reserved by Sony. No idea.
}

func (c *Camera) StartLiveview() (lv *Liveview, err error) {
	resp, err := c.newRequest(endpoints.Camera, "startLiveview").Do()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(resp.Result) > 0 {
		lv = &Liveview{}
		err = json.Unmarshal(resp.Result[0], &lv.URL)
	}

	lvResp, err := http.Get(lv.URL)
	if err != nil {
		return
	}
	lv.Stream = bufio.NewReader(lvResp.Body)
	lv.HTTPResponse = lvResp

	return
}

func (c *Camera) StartLiveviewWithSize(size LiveviewSize) (lv *Liveview, err error) {
	resp, err := c.newRequest(endpoints.Camera, "startLiveviewWithSize", size).Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		lv = &Liveview{}
		err = json.Unmarshal(resp.Result[0], &lv.URL)
	}

	lvResp, err := http.Get(lv.URL)
	if err != nil {
		return
	}
	lv.Stream = bufio.NewReader(lvResp.Body)

	return
}

func (c *Camera) StopLiveview() (err error) {
	_, err = c.newRequest(endpoints.Camera, "stopLiveview").Do()
	return
}

func (lv *Liveview) Decode(out chan *LiveviewPayload) {
	var lastTimestamp int64
	var frames = 0
	for {
		payload, err := lv.decodePayload()
		if err != nil {
			log.Println(err)
			continue
		}
		frames++
		now := time.Now().UnixNano() / 1e6
		if msElapsed := now - lastTimestamp; msElapsed >= 10000 {
			log.Printf("%2.2ffps buffered: %d\n", float64(frames)/(float64(msElapsed)/1000), len(out))
			frames = 0
			lastTimestamp = now
		}

		// Don't let the buffer get full
		if len(out) == cap(out) {
			<-out
		}

		out <- &payload
	}
}

// Stop ends the current livestream
func (lv *Liveview) Stop() {
	lv.HTTPResponse.Body.Close()
	lv.Camera.StopLiveview()
}
