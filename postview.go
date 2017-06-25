package sonycrapi

import "encoding/json"

// PostViewSize defines a supported Post View Image size, as defined by the
// Sony SRC 2.40 API reference. Instances of PostViewSize can be found in the
// PostViewSizes variable.
type PostViewSize string

// PostViewSizes defined the possible size options for postview images
// returned by the camera.
var PostViewSizes = struct {
	Original PostViewSize // The origianl image size
	TwoM     PostViewSize // A device spcific size, smaller than the original. Often 2 megapixels.
}{"Original", "2M"}

// SetPostviewImageSize sets the Post View image size for the camera:
//
// The possible options are:
// "2M" - a smaller preview, usually 2Megpixels in size, sometimes not - camera dependant
// "Original" - the size of the image taken
func (c *Camera) SetPostviewImageSize(size PostViewSize) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setPostviewImageSize", size).Do()
	return
}

// GetPostviewImageSize obtains the current Post View Image size from the camera:
//
// The possible options are:
// "2M" - a smaller preview, usually 2Megpixels in size, sometimes not - camera dependant
// "Original" - the size of the image taken
func (c *Camera) GetPostviewImageSize() (size string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getPostviewImageSize").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &size)
	}

	return
}

// GetSupportedPostviewImageSize obtains the supported Post View Image sizes from the camera
func (c *Camera) GetSupportedPostviewImageSize() (sizes []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getSupportedPostviewImageSize").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &sizes)
	}

	return
}

// GetAvailablePostviewImageSize obtains the current and available Post View Image sizes
// from the camera
func (c *Camera) GetAvailablePostviewImageSize() (current string, available []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getAvailablePostviewImageSize").Do()
	if err != nil {
		return
	}

	if len(resp.Result) >= 1 {
		// Current size
		if err := json.Unmarshal(resp.Result[0], &current); err != nil {
			return current, available, err
		}

		// Available sizes
		if err := json.Unmarshal(resp.Result[1], &available); err != nil {
			return current, available, err
		}
	}

	return
}
