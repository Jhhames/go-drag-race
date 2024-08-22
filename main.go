package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	gameUser := SetupUser()

	fmt.Println("User selected:", gameUser.Name)

	fmt.Println("..........................................................................................")

	game := SetupGame()

	gameUser.Budget = game.Budget

	fmt.Println("Game joined:", game.Id)

	fmt.Println("..........................................................................................")

	if gameUser.Car == (car{}) {
		car := SetupCar()

		gameUser.Car = car

		game = UpdateUser(gameUser, game)

		fmt.Println("..........................................................................................")
	}

	game.StartGame()
}

func SetupUser() User {
	// setup user here first
	var gameUser User

	userSetupChoice := multiChoiceSelect("Hello Welcome back", []string{"Continue as a user", "Add new user"})

	if userSetupChoice == "Continue as a user" {
		userStore, usterr := ReadFromFile("user.txt")

		if usterr != nil {
			fmt.Println("No user exits please create a user.")

			os.Exit(1)
		}

		// read users from file
		userSlice, err := ConvertStringToStructArrayUser(userStore)

		if err != nil {
			fmt.Println("Unable to load user")
			os.Exit(1)
		}

		var userNames []string

		for _, user := range userSlice {
			userNames = append(userNames, user.Name)
		}

		selectUser := multiChoiceSelect("Select a user", userNames)

		for _, user := range userSlice {
			if user.Name == selectUser {
				gameUser = user
				break
			}
		}
	}

	if userSetupChoice == "Add new user" {
		// prompt for user name
		gameUser.Name = oneLinePrompt("Enter user's name")
		gameUser.Id = "ID" + floatToString(getRandom())

		existingUsers, _ := ReadFromFile("user.txt")

		eUsersSlice, _ := ConvertStringToStructArrayUser(existingUsers)
		userSlice := []User{gameUser}

		strString, convertErr := ConvertStructSliceToStringUser(append(eUsersSlice, userSlice...))

		if convertErr != nil {
			fmt.Println("Unable to convert to string", convertErr)
			os.Exit(1)
		}

		err := SaveToFile("user.txt", strString)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Unable to start a game")
			os.Exit(1)
		}
	}

	// return user added
	return gameUser
}

func SetupGame() Game {
	var game Game

	gameChoice := multiChoiceSelect("What would you like to do", []string{"Start a game", "Join a game", "End game"})

	if gameChoice == "End game" {
		err := os.Remove("game.txt")

		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				fmt.Println("unable to end game, no current games in session")
				os.Exit(1)

			}
			fmt.Println("unable to end game", err)
			os.Exit(1)

		}

		fmt.Println("Game ended.")
		os.Exit(1)
	}

	if gameChoice == "Start a game" {
		gameBudget := strToFloat(oneLinePrompt("Enter game budget (number)"))

		game = Game{
			Id:     "ID" + floatToString(getRandom()),
			Budget: gameBudget,
		}
		sliceArr := []Game{game}

		strString, convertErr := ConvertStructSliceToString(sliceArr)

		if convertErr != nil {
			fmt.Println("Unable to convert to string", convertErr)
			os.Exit(1)
		}

		err := SaveToFile("game.txt", strString)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Unable to start a game")
			os.Exit(1)
		}

		fmt.Println("Game created with ID", game.Id)
		fmt.Println("Game created with Budget", game.Budget)
	}

	if gameChoice == "Join a game" {

		gmStr, gmreadErr := ReadFromFile("game.txt")

		if gmreadErr != nil {

			fmt.Println("No Existing game to join")
			os.Exit(1)
		}
		gmSlice, err := ConvertStringToStructArray(gmStr)

		if err != nil {
			fmt.Println("Error converting to struct array", err)
			os.Exit(1)
		}

		game = gmSlice[0]

		fmt.Println("Current available game:", game.Id)
		fmt.Println("Game budget", game.Budget)

		input := oneLinePrompt("Enter Game Id to join game")

		if input != game.Id {
			fmt.Println("Invalid ID")
			// fmt.Println("Setup your user")
			// prompt for name
			//
			os.Exit(1)
		}
	}

	return game
}

func SetupCar() car {

	// budgetCopy := budget
	// const dtMap = map[string]float64 {
	// 	"AWD": 1500,
	// 	"RWD": 1000,
	// }

	var car car

	// enter horsepower
	car = car.configHorsepower()

	// enter torque
	car = car.configTorque()

	// select drive train
	car = car.selectDriveTrain()

	// select chasis weight
	car = car.configWeight()

	car.build()
	// build

	return car
}

func UpdateUser(u User, g Game) Game {
	// find user
	// var findUser User;
	userStore, usterr := ReadFromFile("user.txt")

	if usterr != nil {
		fmt.Println("No user exits please create a user.")

		os.Exit(1)
	}

	// read users from file
	userSlice, err := ConvertStringToStructArrayUser(userStore)

	if err != nil {
		fmt.Println("Unable to load user")
		os.Exit(1)
	}

	for index, user := range userSlice {
		if user.Name == u.Name {
			userSlice[index] = u
			break
		}
	}

	usersStr, _ := ConvertStructSliceToStringUser(userSlice)

	saveUserErr := SaveToFile("user.txt", usersStr)

	if saveUserErr != nil {
		fmt.Println("unable to save users")
	}

	gameStore, gmErr := ReadFromFile("game.txt")

	if gmErr != nil {
		fmt.Println("No games exits please create a game.")

		os.Exit(1)
	}

	// read users from file
	gmSlice, gmSlcErr := ConvertStringToStructArray(gameStore)

	if gmSlcErr != nil {
		fmt.Println("Unable to load games")
		os.Exit(1)
	}

	gmSlice[0].Users = userSlice

	strString, convertErr := ConvertStructSliceToString(gmSlice)

	if convertErr != nil {
		fmt.Println("Unable to convert to string", convertErr)
		os.Exit(1)
	}

	saveGmErr := SaveToFile("game.txt", strString)
	if saveGmErr != nil {
		fmt.Println("unable to save game")
	}

	return gmSlice[0]
}
