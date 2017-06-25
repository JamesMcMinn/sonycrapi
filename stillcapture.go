package sonycrapi

import "encoding/json"

// TakePicture asks the camera to take a still picture It may return an error if the
// camera is not in a ready state, or error 40403 of the camera is taking a long exposue.
// In this case, the client should call AwaitTakePicutre and await the full exposure.
func (c *Camera) TakePicture() (urls []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "actTakePicture").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &urls)
	}

	return
}

// AwaitTakePicture waits while the camera is taking the picture and returns the picture URL
func (c *Camera) AwaitTakePicture() (urls []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "awaitTakePicture").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &urls)
	}

	return
}
