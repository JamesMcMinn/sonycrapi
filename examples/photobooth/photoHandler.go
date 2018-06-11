package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type photoHandler struct {
}

func (ph photoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	photoName := imageName(r.RequestURI)

	if _, err := os.Stat("./photos/" + photoName); os.IsNotExist(err) {
		cLoc := "http://192.168.122.1:8080/postview/memory/DCIM/101MSDCF/" + photoName + "?size=Origin"
		err := downloadPhoto(cLoc)
		if err != nil {
			log.Println("Fuck!")
		}
	}

	log.Println(photoName)
	f, err := os.Open("./photos/" + photoName)
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(f)
	jpg, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(jpg)
}
