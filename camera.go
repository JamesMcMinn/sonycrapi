package sonycrapi

import (
	"encoding/json"
	"log"
)

type Camera struct {
	ActionListURL string

	requestID int

	// API Versions 1.0
	// 0
	AvailableAPIList []string `json:"names"`

	// 1
	Status string `json:"cameraStatus"`

	// 2
	Zoom struct {
		Position           int `json:"zoomPosition"`
		NumberBox          int `json:"zoomNumberBox"`
		CurrentBox         int `json:"zoomIndexCurrentBox"`
		PositionCurrentBox int `json:"zoomPositionCurrentBox"`
	}

	// 3
	LiveviewStatus bool `json:"liveviewStatus"`

	// 4
	LiveviewOrientation string `json:"liveviewOrientation"`

	// 5
	TakePictureURLs []struct {
		URL []string `json:"takePictureUrl"`
	}

	// State 6-9 reserved by sony.

	// 10
	Storage []struct {
		ID                       string `json:"storageID"`
		RecordTarget             bool   `json:"recordTarget"`
		NumberOfRecordableImages int    `json:"numberOfRecordableImages"`
		RecordableTime           int    `json:"recordableTime"` // minutes
		Description              string `json:"storageDescription"`
	}

	// 11
	BeepMode struct {
		Current    string   `json:"currentBeepMode"`
		Candidates []string `json:"beepModeCandidates"`
	}

	// 12
	Function struct {
		Current    string   `json:"currentCameraFunction"`
		Candidates []string `json:cameraFunctionCandidates`
	}

	// 13
	MovieQuality struct {
		Current    string   `json:"currentMovieQuality"`
		Candidates []string `json:"movieQualityCandidates"`
	}

	// 14
	StillSize struct {
		CheckAvailability bool   `json:"checkAvailability"`
		CurrentAspect     string `json:"currentAspect"`
		CurrentSize       string `json:"currentSize"`
	}

	// 15
	FunctionResult string `json:"cameraFunctionResult"`

	// 16
	SteadyMode struct {
		Current    string   `json:"currentSteadyMode"`
		Candidates []string `json:"steadyModeCandidates"`
	}

	// 17
	ViewAngle struct {
		Current    string   `json:"currentViewAngle"`
		Candidates []string `json:"viewAngleCandidates"`
	}

	// 18
	ExposureMode struct {
		Current    string   `json:"currentExposureMode"`
		Candidates []string `json:"exposureModeCandidates"`
	}

	// 19
	PostviewImageSize struct {
		Current    string   `json:"currentPostviewImageSize"`
		Candidates []string `json:"postviewImageSizeCandidates"`
	}

	// 20
	SelfTimer struct {
		Current    int   `json:"currentSelfTimer"`
		Candidates []int `json:"selfTimerCandidates"`
	}

	// 21
	ShootMode struct {
		Current    string   `json:"currentShootMode"`
		Candidates []string `json:"shootModeCandidates"`
	}

	// 22 - 24 Reserved

	// 25
	ExposureCompensation struct {
		Current   int `json:"currentExposureCompensation"`
		Max       int `json:"maxExposureCompensation"`
		Min       int `json:"minExposureCompensation"`
		StepIndex int `json:"stepIndexOfExposureCompensation"`
	}

	// 26
	FlashMode struct {
		Current    string   `json:"currentFlashMode"`
		Candidates []string `json:"flashModeCandidates"`
	}

	// 27
	FNumber struct {
		Current    string   `json:"currentFNumber"`
		Candidates []string `json:"fNumberCandidates"`
	}

	// 28
	FocusMode struct {
		Current    string   `json:"currentFocusMode"`
		Candidates []string `json:"focusModeCandidates"`
	}

	// 29
	ISO struct {
		Current    string   `json:"currentIsoSpeedRate"`
		Candidates []string `json:"isoSpeedRateCandidates"`
	}

	// 30 reserved by sony

	// 31
	IsShifted bool `json:"isShifted"`

	// 32
	ShutterSpeed struct {
		Current    string   `json:"currentShutterSpeed"`
		Candidates []string `json:"shutterSpeedCandidates"`
	}

	// 33
	WhiteBalance struct {
		CheckAvailability bool   `json:"checkAvailability"`
		Current           string `json:"currentWhiteBalanceMode"`
		ColorTemperature  int    `json:"currentColorTemperature"`
	}

	// 34
	TouchAFPosition struct {
		Set                bool      `json:"currentSet"`
		CurrentCoordinates []float64 `json:"currentTouchCoordinates"`
	}

	// Version 1.1

	// 35
	FocusState string `json:"focusStatus"`

	// Version 1.2

	// 36
	ZoomMode struct {
		Current    string   `json:"zoom"`
		Candidates []string `json:"candidate"`
	}

	// 37
	StillQuality struct {
		Current    string   `json:"stillQuality"`
		Candidates []string `json:"candidate"`
	}

	// 38
	ContinuousShootingMode struct {
		Current    string   `json:"contShootingMode"`
		Candidates []string `json:"candidate"`
	}

	// 39
	ContinuousShootingSpeed struct {
		Current    string   `json:"contShootingSpeed"`
		Candidates []string `json:"candidate"`
	}

	// 40
	ContinuousShootingURL []struct {
		PostviewURL  string `json:"postviewUrl"`
		ThumbnailURL string `json:"thumbnailUrl"`
	}

	// 41
	FlipSetting struct {
		Current    string   `json:"flip"`
		Candidates []string `json:"candidate"`
	}

	// 42
	SceneSelection struct {
		Current    string   `json:"scene"`
		Candidates []string `json:"candidate"`
	}

	// 43
	InvtervalTime struct {
		Current    string   `json:`
		Candidates []string `json:`
	}

	// 44
	Color struct {
		Current    string   `json:"colorSetting"`
		Candidates []string `json:"candidate"`
	}

	// 45
	MovieFileFormat struct {
		Current    string   `json:"movieFileFormat"`
		Candidates []string `json:"candidate"`
	}

	// 46 - 51 reserved

	// 52
	IRRemoteSetting struct {
		Current    string   `json:"infraredRemoteControl"`
		Candidates []string `json:"candidate"`
	}

	// 53
	TVColorSystem struct {
		Current    string   `json:"tvColorSystem"`
		Candidates []string `json:"candidate"`
	}

	// 54
	TrackingFocusStauts string `json:"trackingFocusStatus"`

	// 55
	TrackingFocusSetting struct {
		Current    string   `json:"trackingFocus"`
		Candidates []string `json:"candidate"`
	}

	// 56
	BatteryInfo []struct {
		ID               string `json:"batteryID"`
		Status           string `json:"status"`
		AdditionalStatus string `json:"additionalStatus"`
		LevelNumber      int    `json:"levelNumer"`
		LevelDenom       int    `json:"levelDenom"`
		Description      string `json:"description"`
	}

	// 57
	RecordingTime int `json:"recordingTime"`

	// 58
	NumberOfshots int `json:"numberOfShots"`

	// 59
	AutoPowerOffTime struct {
		Current    int   `json:"autoPowerOff"`
		Candidates []int `json:"candidate"`
	}

	// Version 1.3

	// 60
	LoopRecordingTime struct {
		Current    string   `json:"loopRecTime"`
		Candidates []string `json:"candidate"`
	}

	// 61
	AudioRecording struct {
		Current    string   `json:"audioRecording"`
		Candidates []string `json:"candidate"`
	}

	// 62
	WindNoiseReduction struct {
		Current    string   `json:"windNoiseReduction"`
		Candidates []string `json:"candidate"`
	}
}

func NewCamera(actionListURL string) *Camera {
	return &Camera{
		ActionListURL: actionListURL,
	}
}

func (c *Camera) Initilize() {

}

func extractNonType(data json.RawMessage) json.RawMessage {
	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		log.Println(err)
	}

	delete(m, "type")

	if n := len(m); n > 1 {
		log.Println("Warning: Expected single value after type removal. Found", n)
	}

	for _, val := range m {
		return val
	}

	// Returning nothing is there was nothing...
	return []byte("\"\"")
}

func (c *Camera) GetEventState() (err error) {
	resp, err := c.newRequest(endpoints.Camera, "getEvent", false).Do()
	if err != nil {
		return
	}

	for i, data := range resp.Result {
		switch i {
		case 0:
			err = json.Unmarshal(extractNonType(data), &c.AvailableAPIList)
		case 1:
			err = json.Unmarshal(extractNonType(data), &c.Status)
		case 2:
			err = json.Unmarshal(data, &c.Zoom)
		case 3:
			err = json.Unmarshal(extractNonType(data), &c.LiveviewStatus)
		case 4:
			err = json.Unmarshal(extractNonType(data), &c.LiveviewOrientation)
		case 5:
		// TODO
		case 10:
			err = json.Unmarshal(data, &c.Storage)
		case 11:
			err = json.Unmarshal(data, &c.BeepMode)
		case 12:
			err = json.Unmarshal(data, &c.Function)
		case 13:
			err = json.Unmarshal(data, &c.MovieQuality)
		case 14:
			err = json.Unmarshal(data, &c.StillSize)
		case 15:
			err = json.Unmarshal(extractNonType(data), &c.FunctionResult)
		case 16:
			err = json.Unmarshal(data, &c.SteadyMode)
		case 17:
			err = json.Unmarshal(data, &c.ViewAngle)
		case 18:
			err = json.Unmarshal(data, &c.ExposureMode)
		case 19:
			err = json.Unmarshal(data, &c.PostviewImageSize)
		case 20:
			err = json.Unmarshal(data, &c.SelfTimer)
		case 21:
			err = json.Unmarshal(data, &c.ShootMode)
		case 22:
			// Reserved
		case 23:
			// Reserved
		case 24:
			// Reserved
		case 25:
			err = json.Unmarshal(data, &c.ExposureCompensation)
		case 26:
			err = json.Unmarshal(data, &c.FlashMode)
		case 27:
			err = json.Unmarshal(data, &c.FNumber)
		case 28:
			err = json.Unmarshal(data, &c.FocusMode)
		case 29:
			err = json.Unmarshal(data, &c.ISO)
		case 30:
			// Reserved
		case 31:
			err = json.Unmarshal(extractNonType(data), &c.IsShifted)
		case 32:
			err = json.Unmarshal(data, &c.ShutterSpeed)
		case 33:
			err = json.Unmarshal(data, &c.WhiteBalance)
		case 34:
			err = json.Unmarshal(data, &c.TouchAFPosition)

		// Version 1.1
		case 35:
			err = json.Unmarshal(extractNonType(data), &c.FocusState)

			// Version 1.2
		case 36:
			err = json.Unmarshal(data, &c.ZoomMode)
		case 37:
			err = json.Unmarshal(data, &c.StillQuality)
		case 38:
			err = json.Unmarshal(data, &c.ContinuousShootingMode)
		case 39:
			err = json.Unmarshal(data, &c.ContinuousShootingSpeed)
		case 40:
			err = json.Unmarshal(data, &c.ContinuousShootingURL)
		case 41:
			err = json.Unmarshal(data, &c.FlipSetting)
		case 42:
			err = json.Unmarshal(data, &c.SceneSelection)
		case 43:
			err = json.Unmarshal(data, &c.InvtervalTime)
		case 44:
			err = json.Unmarshal(data, &c.Color)
		case 45:
			err = json.Unmarshal(data, &c.MovieFileFormat)
		case 46:
			//Reserved
		case 47:
			//Reserved
		case 48:
			//Reserved
		case 49:
			//Reserved
		case 50:
			//Reserved
		case 51:
			//Reserved
		case 52:
			err = json.Unmarshal(data, &c.IRRemoteSetting)
		case 53:
			err = json.Unmarshal(data, &c.TVColorSystem)
		case 54:
			err = json.Unmarshal(extractNonType(data), &c.TrackingFocusStauts)
		case 55:
			err = json.Unmarshal(data, &c.TrackingFocusSetting)
		case 56:
			err = json.Unmarshal(data, &c.BatteryInfo)
		case 57:
			err = json.Unmarshal(extractNonType(data), &c.RecordingTime)
		case 58:
			err = json.Unmarshal(extractNonType(data), &c.NumberOfshots)
		case 59:
			err = json.Unmarshal(data, &c.AutoPowerOffTime)

			//Version 1.3
		case 60:
			err = json.Unmarshal(data, &c.LoopRecordingTime)
		case 61:
			err = json.Unmarshal(data, &c.AudioRecording)
		case 62:
			err = json.Unmarshal(data, &c.WindNoiseReduction)
		}

		if err != nil {
			log.Println(i, err)
		}
	}

	return
}

// Fuck this file in particular.
