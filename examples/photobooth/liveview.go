package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jamesmcminn/sonycrapi"
)

// Latency is more important than refresh rate for the camera - so smalls buffers for each connection.
const frameBufferSize = 3

var frames = make(chan *sonycrapi.LiveviewPayload, 1)

var (
	liveviewConnections = map[*http.Request](chan []byte){}
	lvConnLock          = &sync.RWMutex{}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  16384,
	WriteBufferSize: 16384,
	CheckOrigin: func(r *http.Request) bool { // Accept any websocket connection. We don't care about origin
		return true
	},
}

func fillFrameBuffers() {
	for {
		frame := <-frames

		b64 := []byte(base64.StdEncoding.EncodeToString(frame.JPEGData))

		lvConnLock.RLock()
		for _, buffer := range liveviewConnections {
			if len(buffer) == cap(buffer) {
				log.Println("Dropping frame.")
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

	buffer := make(chan []byte, 5)
	lvConnLock.Lock()
	liveviewConnections[r] = buffer
	lvConnLock.Unlock()
	log.Println("New stream connection... Total:", len(liveviewConnections))

	for {
		frame := <-buffer
		err := conn.WriteMessage(websocket.BinaryMessage, frame)
		if err != nil {
			break
		}
	}

	lvConnLock.Lock()
	delete(liveviewConnections, r)
	lvConnLock.Unlock()
	log.Println("Live veiw connection ended... Total:", len(liveviewConnections))
}
