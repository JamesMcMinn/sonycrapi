package sonycrapi

import (
	"fmt"
	"testing"
)

func TestContinuousShooting(t *testing.T) {
	camera := NewCamera(cameraURL)

	av, err := camera.GetSupportedShootMode()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(av)

	camera.SetShootMode(ShootModes.Still)

	err = camera.SetContinuousShootingMode(ContinuousShootingModes.Continuous)
	if err != nil {
		t.Error(err)
	}

	// current, avail, err := camera.GetAvailableContinuousShootingMode()
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println("Current:", current, "Avail:", avail)

	// err = camera.SetContinuousShootingSpeed(ContinuousShootingSpeeds.TwoFPSx5)
	// if err != nil {
	// 	t.Error(err)
	// }

	// err = camera.StartContinuousShooting()
	// if err != nil {
	// 	t.Error(err)
	// }

	// err = camera.StopContinuousShooting()
	// if err != nil {
	// 	t.Error(err)
	// }

}
