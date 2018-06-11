package sonycrapi

import "encoding/json"

// StillQuality defines all possible quality levels for still photos
type continuousShootingMode string

var ContinuousShootingModes = struct {
	Single      continuousShootingMode
	Continuous  continuousShootingMode
	SpdPriority continuousShootingMode
	Burst       continuousShootingMode
	MotionShot  continuousShootingMode
}{"Single", "Continuous", "Spd Priority Cont.", "Burst", "MotionShot"}

type continuousShootingSpeed string

var ContinuousShootingSpeeds = struct {
	Hi         continuousShootingSpeed
	Low        continuousShootingSpeed
	TenFPSx1   continuousShootingSpeed
	EightFPSx1 continuousShootingSpeed
	FiveFPSx2  continuousShootingSpeed
	TwoFPSx5   continuousShootingSpeed
}{"Hi", "Low", "10fps 1sec", "8fps 1sec", "5fps 2sec", "2fps 5sec"}

func (c *Camera) SetContinuousShootingMode(mode continuousShootingMode) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setContShootingMode",
		map[string]continuousShootingMode{"contShootingMode": mode}).Do()
	return
}

func (c *Camera) GetContinuousShootingMode() (mode continuousShootingMode, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getContShootingMode").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {

		csm := struct {
			ContShootingMode continuousShootingMode `json:"contShootingMode"`
		}{}
		err = json.Unmarshal(resp.Result[0], &csm)
		if err != nil {
			return
		}

		mode = csm.ContShootingMode
	}

	return
}

func (c *Camera) GetSupportedContinuousShootingMode() (modes []continuousShootingMode, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getSupportedContShootingMode").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		candiates := struct {
			Candidate []continuousShootingMode `json:"candidate"`
		}{}
		err = json.Unmarshal(resp.Result[0], &candiates)
		if err != nil {
			return
		}

		modes = candiates.Candidate
	}

	return
}

func (c *Camera) GetAvailableContinuousShootingMode() (current continuousShootingMode, available []continuousShootingMode, err error) {
	resp, err := c.newRequest(endpoints.Camera, "getAvailableContShootingMode").Do()
	if err != nil {
		return
	}

	if len(resp.Result) > 0 {
		avbMds := struct {
			Current   continuousShootingMode   `json:"contShootingMode"`
			Candidate []continuousShootingMode `json:"candidate"`
		}{}
		err = json.Unmarshal(resp.Result[0], &avbMds)
		if err != nil {
			return
		}

		current, available = avbMds.Current, avbMds.Candidate
	}

	return
}

// SetContinuousShootingSpeed sets the shotting speed of the camera
func (c *Camera) SetContinuousShootingSpeed(speed continuousShootingSpeed) (err error) {
	_, err = c.newRequest(endpoints.Camera, "setContShootingSpeed",
		map[string]continuousShootingSpeed{"contShootingSpeed": speed}).Do()
	return
}

// StartContinuousShooting informs the camera to begin continuous shooting
func (c *Camera) StartContinuousShooting() (err error) {
	_, err = c.newRequest(endpoints.Camera, "startContShooting").Do()
	return
}

// StopContinuousShooting informs the camera to stop continuous shooting
func (c *Camera) StopContinuousShooting() (err error) {
	_, err = c.newRequest(endpoints.Camera, "stopContShooting").Do()
	return
}
