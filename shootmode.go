package sonycrapi

import "encoding/json"

// ShootMode defines a particualr shooting mode for the camera.
// Instances shootmode can be found in the ShootModes variable for use
// in SetShootMode()
type ShootMode string

// ShootModes defines, at the time of writing, all possible shoot
// modes supported by the camera API.
var ShootModes = struct {
	Still         ShootMode // Still image shoot mode
	Movie         ShootMode // Movie shoot mode
	Audio         ShootMode // Audio shoot mode
	IntervalStill ShootMode // Interval still shoot mode
	LoopRec       ShootMode // Loop recording shoot mode
}{"still", "movie", "audio", "intervalstill", "looprec"}

// SetShootMode sets the camera to the specified shoot mood.
// All possible shoot modes are defined in the ShootMode sturct.
func (c *Camera) SetShootMode(mode ShootMode) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setShootMode", mode).Do()
	return
}

// GetShootMode returns the current shoot mode of the camera.
func (c *Camera) GetShootMode() (mode string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getShootMode").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &mode)
	}

	return
}

// GetSupportedShootMode obtains the supported shooting modes for the camera
func (c *Camera) GetSupportedShootMode() (modes []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getSupportedShootMode").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		err = json.Unmarshal(resp.Result[0], &modes)
	}

	return
}

// GetAvailableShootMode obtains the current and available shoot mode for
// the camera
func (c *Camera) GetAvailableShootMode() (current string, available []string, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getAvailableShootMode").Do()
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
