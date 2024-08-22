package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func oneLinePrompt(text string) string {
	fmt.Println(text, ":")
	reader := bufio.NewReader(os.Stdin)
	prompt, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Remove the newline character from the prompt
	prompt = prompt[:len(prompt)-1]

	// fmt.Println("Prompt:", prompt)
	return prompt
}

func multiChoiceSelect(text string, options []string) string {
	// options := []string{"Option 1", "Option 2", "Option 3"}

	fmt.Println(text, ":")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	choice, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid choice:", input[:len(input)-1])
		os.Exit(1)
	}

	if choice < 1 || choice > len(options) {
		fmt.Println("Invalid choice:", choice)
		os.Exit(1)

	}

	selectedOption := options[choice-1]

	return selectedOption
}

func strToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		os.Exit(1)
	}

	return f
}

func ConvertStructSliceToString(game []Game) (string, error) {
	data, err := json.Marshal(game)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
func ConvertStructSliceToStringUser(user []User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ConvertStringToStructArray(str string) ([]Game, error) {
	var games []Game
	err := json.Unmarshal([]byte(str), &games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func ConvertStringToStructArrayUser(str string) ([]User, error) {
	var users []User
	err := json.Unmarshal([]byte(str), &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getRandom() int {
	rand.Seed(time.Now().UnixNano())

	min := 100
	max := 200

	return rand.Intn(max-min+1) + min
}

func floatToString(num int) string {
	return strconv.Itoa(num)
}