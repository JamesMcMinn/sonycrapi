package sonycrapi

// StillQuality defines all possible quality levels for still photos
var StillQuality = struct {
	RAWJPG   string
	Fine     string
	Standard string
}{"RAW+JPEG", "Fine", "Standard"}

// SetStillQuality sets the quality for still photos taken by the camera.
// The possible quality settings are defined in the StillQuality struct.
func (c *Camera) SetStillQuality(stillQuality string) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setStillQuality",
		map[string]string{"stillQuality": stillQuality}).Do()
	return
}
