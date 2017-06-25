package sonycrapi

import (
	"fmt"
	"testing"
)

func TestTakePicture(t *testing.T) {
	camera := NewCamera(cameraURL)
	// camera.SetPostviewImageSize(PostViewSizes.TwoM)
	urls, err := camera.TakePicture()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(urls)
}
