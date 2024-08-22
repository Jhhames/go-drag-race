package main

import (
	"fmt"
	"os"
)

type User struct {
	Id string
	Name string
	Car car
	Budget float64
}

type Game struct {
	Id     string
	Budget float64
	Users [] User
}

func (g Game)StartGame() {
	startChoice := multiChoiceSelect("Your Car is setup, do you want to start game?", []string{"Yes", "No"})

	if startChoice != "Yes" {
		fmt.Println("Ready to start game when you are")
		os.Exit(1)
	}

	if len(g.Users) < 2 {
		fmt.Println("Can't start game with less than 2 users")
		os.Exit(0)
	}


	for _, user := range g.Users {
		if user.Car == (car{}) {
			fmt.Println(fmt.Sprintf("%s's car is not setup yet", user.Name))
			os.Exit(0)
		}
	}

	for _, user := range g.Users {
		qtrMile := user.Car.quarterMileTime()

		m := fmt.Sprintf("%s finished the quarter mile in %f seconds", user.Name, qtrMile)

		fmt.Println(m)
	}

	fmt.Println("Game ended thanks for playing")
}