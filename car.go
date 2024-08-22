package main

import (
	"fmt"
	"math"
	"os"
)

type car struct {
	Horsepower float64
	Torque     float64
	DriveTrain string
	Weight     float64
	UserID string
	Build bool
}

func (c car) configHorsepower() car {
	c.Horsepower = strToFloat(oneLinePrompt("Enter Car Horsepower"))

	return c
}
func (c car) configTorque() car {
	c.Torque = strToFloat(oneLinePrompt("Enter Car Torque"))

	return c
}
func (c car) configWeight() car {
	c.Weight = strToFloat(oneLinePrompt("Enter Car chasis weight"))

	return c
}

func (c car) selectDriveTrain() car {
	c.DriveTrain = multiChoiceSelect("Select Drive Train", []string{"AWD", "RWD", "FWD"})

	return c
}

func (c car) build() car {
	if c.Horsepower < 0 || c.Horsepower > 1000 {
		fmt.Println("Invalid horspower unable to build")
		os.Exit(1)
	}
	if c.Torque < 0 || c.Torque > 1000 {
		fmt.Println("Invalid torque unable to build")
		os.Exit(1)
	}

	if c.Weight < 1400 || c.Weight > 4000 {
		fmt.Println("Invalid weight unable to build")
		os.Exit(1)
	}

	if c.DriveTrain == "" {
		fmt.Println("Invalid horspower unable to build Car")
		os.Exit(1)
	}

	c.Build = true

	return c
}

func (c car) quarterMileTime() float64 {

	// Convert the horsepower and torque to metric units
	horsepower := c.Horsepower * 0.7457
	torque := c.Torque * 0.113
	weight := c.Weight

	// Calculate the power of the car
	power := horsepower * torque / 5252

	// Calculate the acceleration of the car
	acceleration := power / weight

	distance := 402.336

	timeCoeff := math.Sqrt(2 * distance / acceleration)

	// Return the time in seconds.
	return timeCoeff / 100
}
