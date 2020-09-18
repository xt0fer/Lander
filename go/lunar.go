/*
 * main.go
 *
 * Copyright 2020 ZipCodeWilmington Kris Younger
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */
package lunar

import (
	"fmt"
	"math/rand"
)

func calculate(height, speed, burn, gravity int) int {
	return (speed + gravity - burn)
}

func windowcleaner(step int) int {
	if step >= 24 {
		fmt.Printf("\nTime\t")
		fmt.Printf("Speed\t\t")
		fmt.Printf("Fuel\t\t")
		fmt.Printf("Height\t\t")
		fmt.Printf("Burn\n")
		fmt.Printf("----\t")
		fmt.Printf("-----\t\t")
		fmt.Printf("----\t\t")
		fmt.Printf("------\t\t")
		fmt.Printf("----\n")
		step = 1
	} else {
		if step < 24 {
			step++
		}
	}
	return step
}

func randomheight() int {
	rnd := rand.New(rand.NewSource(99))
	r := rnd.Int()
	return (r%15000 + 4000)
}

type Vehicle struct {
	Height     int /* The height of the spaceship. */
	Speed      int /* The speed of the spaceship. */
	Burn       int /* The fuel which gets burned this step */
	Tensec     int /* The time the flight is running for. (in ten second steps) */
	Fuel       int /* The fuel you have left. (kilogram) */
	PrevHeight int /* The previous height to compare with actual. (for the colored digits) */
	Step       int /* Counts the steps passed since last output of the column names */

}

const Gravity = 100 /* The rate in which the spaceship descents in free fall (in ten seconds) */

const (
	version   = "1.1" /* The Version of the program */
	dead      = "\nThere were no survivors.\n\n"
	crashed   = "\nThe Spaceship crashed. Good luck getting back home.\n\n"
	success   = "\nYou made it! Good job!\n\n"
	emptyfuel = "\nThere is no fuel left. You're floating around like Wheatley.\n\n"
)

const (
	redFormatNumber   = "\x1b[31m%d\x1b[0m\t\t"
	greenFormatNumber = "\x1b[32m%d\x1b[0m\t\t"
)

func printHeader() {
	fmt.Printf("\nLunar Lander - Version %s\n", version)
	fmt.Printf("This is a computer simulation of an Apollo lunar landing capsule.\n")
	fmt.Printf("The on-board computer has failed so you have to land the capsule manually.\n")
	fmt.Printf("Set burn rate of retro rockets to any value between 0 (free fall) and 200\n")
	fmt.Printf("(maximum burn) kilo per second. Set burn rate every 10 seconds.\n") /* That's why we have to go with 10 second-steps. */
	fmt.Printf("You must land at a speed of 2 or 1. Good Luck!\n\n")
	fmt.Printf("\nTime\t")
	fmt.Printf("Speed\t\t")
	fmt.Printf("Fuel\t\t")
	fmt.Printf("Height\t\t")
	fmt.Printf("Burn\n")
	fmt.Printf("----\t")
	fmt.Printf("-----\t\t")
	fmt.Printf("----\t\t")
	fmt.Printf("------\t\t")
	fmt.Printf("----\n")
}

func getBurnRate() int {
	burn := 0
	// Do very simple input validation.
	for {
		_, err := fmt.Scanf("%d", &burn)
		if err != nil {
			fmt.Printf("Burn rate needs to be a number.\n")
			continue
		}
		if burn < 0 || burn > 200 { /* If there is a wrong entry */
			fmt.Printf("The burn rate rate must be between 0 and 200.\n> ")
			continue
		} else {
			break
		}
	}
	return burn
}

func (vehicle *Vehicle) checkVehicleStatus() string {
	s := ""
	if vehicle.Height <= 0 {
		if vehicle.Speed > 10 {
			s = fmt.Sprintf("%s", dead)
		}
		if vehicle.Speed < 10 && vehicle.Speed > 3 {
			s = fmt.Sprintf("%s", crashed)
		}

		if vehicle.Speed < 3 {
			s = fmt.Sprintf("%s", success)
		}
	} else {
		if vehicle.Height > 0 {
			s = fmt.Sprintf("%s", emptyfuel)
		}
	}
	return s
}

func (vehicle *Vehicle) adjustForBurn() {
	vehicle.PrevHeight = vehicle.Height
	vehicle.Speed = calculate(vehicle.Height, vehicle.Speed, vehicle.Burn, Gravity)
	vehicle.Height = vehicle.Height - vehicle.Speed
	vehicle.Fuel = vehicle.Fuel - vehicle.Burn
}

func (vehicle *Vehicle) stillFlying() bool {
	return vehicle.Height > 0
}

func (vehicle *Vehicle) getStatusLine() string {
	s := ""
	s = s + fmt.Sprintf("%d0\t", vehicle.Tensec)
	s = s + fmt.Sprintf("%d\t\t", vehicle.Speed)
	s = s + fmt.Sprintf("%d\t\t", vehicle.Fuel)

	if vehicle.Height < vehicle.PrevHeight {
		s = s + fmt.Sprintf(redFormatNumber, vehicle.Height)
	}
	if vehicle.Height == vehicle.PrevHeight {
		s = s + fmt.Sprintf("%d\t\t", vehicle.Height)
	}

	if vehicle.Height > vehicle.PrevHeight {
		s = s + fmt.Sprintf(greenFormatNumber, vehicle.Height)
	}
	return s
}

func RunSimulation() {
	status := ""

	/* Set initial height, time, fuel, burn, prevheight, step and speed according to difficulty. */
	h := randomheight()
	vehicle := &Vehicle{
		Height:     h,
		Speed:      1000,
		Fuel:       12000,
		Tensec:     0,
		Burn:       0,
		PrevHeight: h,
		Step:       1,
	}

	printHeader()

	for vehicle.stillFlying() {

		vehicle.Step = windowcleaner(vehicle.Step)

		status = vehicle.getStatusLine()
		fmt.Printf("%s", status)

		vehicle.Burn = getBurnRate()

		vehicle.adjustForBurn()

		if vehicle.Fuel <= 0 {
			break
		}

		vehicle.Tensec++

	}

	status = vehicle.checkVehicleStatus()
	fmt.Printf("%s", status)

	return
}
