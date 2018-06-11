package sonycrapi

// StillQuality defines all possible quality levels for still photos
type stillQuality string

// StillQualities defines the possible qualities that can be passed
// to the camera.
var StillQualities = struct {
	RAWJPG   stillQuality
	Fine     stillQuality
	Standard stillQuality
}{"RAW+JPEG", "Fine", "Standard"}

// SetStillQuality sets the quality for still photos taken by the camera.
// The possible quality settings are defined in the StillQuality struct.
func (c *Camera) SetStillQuality(quality stillQuality) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setStillQuality",
		map[string]stillQuality{"stillQuality": quality}).Do()
	return
}
