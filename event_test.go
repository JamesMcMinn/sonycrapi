package sonycrapi

import "testing"

func TestGetEventState(t *testing.T) {
	camera := NewCamera(cameraURL)
	camera.GetEventState()
}
