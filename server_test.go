package sonycrapi

import (
	"fmt"
	"log"
	"testing"
)

func TestGetAvailableAPIList(t *testing.T) {
	camera := NewCamera(cameraURL)

	av, err := camera.GetAvailableAPIList()
	if err != nil {
		t.Error(err)
	}

	if len(av) == 0 {
		t.Error("Got no available APIs from the camera.")
	}
}

func TestGetApplicationInfo(t *testing.T) {
	camera := NewCamera(cameraURL)

	name, version, err := camera.GetApplicationInfo()
	if err != nil {
		t.Error(err)
	}

	if name == "" || version == "" {
		t.Errorf("Expected a name and a version, instead got \"%s\" and \"%s\".", name, version)
	}

	fmt.Println(name, version)
}

func TestGetVersions(t *testing.T) {
	camera := NewCamera(cameraURL)

	versions, err := camera.GetVersions(endpoints.Camera)
	if err != nil {
		t.Error(err)
	}

	if len(versions) == 0 {
		t.Error("Got no versions back for camera.")
	}

	log.Println("Camera version:", versions)
}

// This fails for me - might actually be an issue with the camera rather than
// my code though...?
func TestGetMethodTypes(t *testing.T) {
	camera := NewCamera(cameraURL)

	types, err := camera.GetMethodTypes(endpoints.Camera, "")
	if err != nil {
		t.Error(err)
	}

	if len(types) == 0 {
		t.Error("Got no method types back for camera.")
	}
}
