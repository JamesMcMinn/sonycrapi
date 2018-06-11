package main

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Latency is more important than refresh rate for the camera - so smalls buffers for each connection.
const frameBufferSize = 3

var (
	liveviewConnections = map[*http.Request](chan []byte){}
	lvConnLock          = &sync.RWMutex{}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // Accept any websocket connection. We don't care about origin
		return true
	},
}

func fillFrameBuffers() {
	for {
		frame := <-diffFrames

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, frame, nil)
		if err != nil {
			log.Println(err)
		}

		b64 := []byte(base64.StdEncoding.EncodeToString(buf.Bytes()))

		lvConnLock.RLock()
		for _, buffer := range liveviewConnections {
			if len(buffer) == cap(buffer) {
				// If the buffer is full, start dropping frames.
				<-buffer
			}
			buffer <- b64
		}
		lvConnLock.RUnlock()
	}
}

func liveviewHTTPHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	buffer := make(chan []byte, 3)
	lvConnLock.Lock()
	liveviewConnections[r] = buffer
	lvConnLock.Unlock()
	log.Println("New stream connection... Total:", len(liveviewConnections))

	for {
		frame := <-buffer
		err := conn.WriteMessage(1, frame)
		if err != nil {
			break
		}
	}

	lvConnLock.Lock()
	delete(liveviewConnections, r)
	lvConnLock.Unlock()
	log.Println("Live veiw connection ended... Total:", len(liveviewConnections))
}
