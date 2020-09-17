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

func help() {
	fmt.Printf("Lunar Lander version 1.1\n")
	fmt.Printf("Made by Kristofer\n\n")
	fmt.Printf("The following arguments are possible (only one):\n")
	fmt.Printf("-d [1/2/3]\tDefine difficulty. 1 is easy 3 is hard.\n")
	fmt.Printf("--info\tShow different intro.")
	fmt.Printf("--help\tPrint this help and exit.\n")
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
	Prevheight int /* The previous height to compare with actual. (for the colored digits) */
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

func RunSimulation() {

	/* Set initial height, time, fuel, burn, prevheight, step and speed according to difficulty. */
	h := randomheight()
	vehicle := &Vehicle{
		Height:     h,
		Speed:      1000,
		Fuel:       12000,
		Tensec:     0,
		Burn:       0,
		Prevheight: h,
		Step:       1,
	}

	printHeader()

	for vehicle.Height > 0 {

		vehicle.Step = windowcleaner(vehicle.Step)

		fmt.Printf("%d0\t", vehicle.Tensec)
		fmt.Printf("%d\t\t", vehicle.Speed)
		fmt.Printf("%d\t\t", vehicle.Fuel)

		if vehicle.Height < vehicle.Prevheight {
			fmt.Printf("\x1b[31m%d\x1b[0m\t\t", vehicle.Height)
		}
		if vehicle.Height == vehicle.Prevheight {
			fmt.Printf("%d\t\t", vehicle.Height)
		}

		if vehicle.Height > vehicle.Prevheight {
			fmt.Printf("\x1b[32m%d\x1b[0m\t\t", vehicle.Height)
		}

		fmt.Scanf("%d", &vehicle.Burn)

		if vehicle.Burn < 0 || vehicle.Burn > 200 { /* If there is a wrong entry */
			fmt.Printf("The burn rate rate must be between 0 and 200.\n")
			continue
		}

		vehicle.Prevheight = vehicle.Height
		vehicle.Speed = calculate(vehicle.Height, vehicle.Speed, vehicle.Burn, Gravity)
		vehicle.Height = vehicle.Height - vehicle.Speed
		vehicle.Fuel = vehicle.Fuel - vehicle.Burn

		if vehicle.Fuel <= 0 {
			break
		}

		vehicle.Tensec++

	}

	if vehicle.Height <= 0 {
		if vehicle.Speed > 10 {
			fmt.Printf("%s", dead)
		}
		if vehicle.Speed < 10 && vehicle.Speed > 3 {
			fmt.Printf("%s", crashed)
		}

		if vehicle.Speed < 3 {
			fmt.Printf("%s", success)
		}
	} else {
		if vehicle.Height > 0 {
			fmt.Printf("%s", emptyfuel)
		}
	}
	return
}
