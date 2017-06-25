package sonycrapi

import "testing"

func TestSetPostviewImageSize(t *testing.T) {
	camera := NewCamera(cameraURL)

	setTo := PostViewSizes.Original
	if err := camera.SetPostviewImageSize(PostViewSizes.Original); err != nil {
		t.Error(err)
	}

	size, err := camera.GetPostviewImageSize()
	if err != nil {
		t.Error(err)
	}

	if size != string(setTo) {
		t.Errorf("Post View Image size: expected %s, got %s", setTo, size)
	}

	if err := camera.SetPostviewImageSize(PostViewSizes.TwoM); err != nil {
		t.Error(err)
	}
}
