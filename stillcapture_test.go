package sonycrapi

import (
	"fmt"
	"testing"
)

func TestTakePicture(t *testing.T) {
	camera := NewCamera(cameraURL)

	// camera.UpdateState(false)
	camera.newRequest(endpoints.Camera, "startRecMode").Do()
	camera.SetShootMode(ShootModes.Still)

	urls, err := camera.TakePicture()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(urls)
}
