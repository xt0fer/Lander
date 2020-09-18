package lunar

import (
	"fmt"
	"testing"
)

// This TestLunar basically lets the lander fall out of the sky.
func TestLunar(t *testing.T) {
	status := ""

	/* Set initial vehicle parameters */
	vehicle := &Vehicle{
		Height:     4000,
		Speed:      1000,
		Fuel:       12000,
		Tensec:     0,
		Burn:       0,
		PrevHeight: 4000,
		Step:       1,
	}
	burns := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	burnIdx := 0

	for vehicle.stillFlying() {

		vehicle.Burn = burns[burnIdx]
		burnIdx++

		vehicle.adjustForBurn()

		if vehicle.outOfFuel() {
			break
		}

		vehicle.Tensec++

	}

	status = vehicle.checkStatus()
	if status == dead {
		fmt.Printf("%s %s", status, "There were no survivors.")
	} else {
		t.Fail()
	}
}
