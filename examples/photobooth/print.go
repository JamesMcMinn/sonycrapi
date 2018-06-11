package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type printHandler struct {
}

func (ph printHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	photoName := imageName(r.RequestURI)
	log.Println(photoName)

	if _, err := os.Stat("./photos/" + photoName); os.IsNotExist(err) {
		cLoc := "http://192.168.122.1:8080/postview/memory/DCIM/101MSDCF/" + photoName + "?size=Origin"
		err := downloadPhoto(cLoc)
		if err != nil {
			log.Println("Fuck!")
			return
		}
	}

	// lp -o fit-to-page -o landscape  preview-test.jpg

	log.Println("Attempting to print...")

	cmd := exec.Command("lp", "-o fit-to-page", "-o landscape", "photos/"+photoName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

}
