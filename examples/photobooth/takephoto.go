package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func takephotoHTTPHandler(w http.ResponseWriter, r *http.Request) {
	urls, err := camera.TakePicture()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(urls)

	photoName := ""
	if len(urls) > 0 {
		cameraURL := urls[0]
		photoName = imageName(cameraURL)
	}

	response := struct {
		Error     error
		PhotoName string
	}{
		err,
		photoName,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func downloadPhoto(url string) (err error) {
	response, e := http.Get(url)
	if e != nil {
		return err
	}

	defer response.Body.Close()

	photoPath := "./photos/" + imageName(url)
	tempName := "./photos/" + fmt.Sprintf("%d", time.Now().UnixNano())
	file, err := os.Create(tempName)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()

	err = os.Rename(tempName, photoPath)
	if err != nil {
		log.Println(err)
	}

	return
}

func imageName(url string) string {
	pos := strings.LastIndex(url, "/")

	if pos < 0 {
		return ""
	}

	name := url[pos+1:]
	if qPos := strings.Index(name, "?"); qPos < 0 {
		return name
	} else {
		return name[:qPos]
	}

}
