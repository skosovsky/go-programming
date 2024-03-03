package main

import "fmt"

type Device struct {
	On    bool
	Ammo  int
	Power int
}

func (d *Device) Shoot() bool {
	if !d.On || d.Ammo == 0 {
		return false
	}

	d.Ammo--
	return true
}

func (d *Device) RideBike() bool {
	if !d.On || d.Power == 0 {
		return false
	}

	d.Power--
	return true
}

func main() {
	testStruct := &Device{false, 5, 5}
	testStruct.Shoot()
	fmt.Print(testStruct)
}
