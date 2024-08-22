package main

import (
	"io/ioutil"

)

func SaveToFile(filename string, data string) error {

	return ioutil.WriteFile(filename, []byte(data), 0666)
}

func ReadFromFile(filename string) (string, error) {
	byteRes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(byteRes), nil
}