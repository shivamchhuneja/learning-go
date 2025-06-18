package helpers

import (
	"errors"
	"fmt"
)

func GetUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func StrUserInput(infoText string) (string, error) {
	var userInput string

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput, nil
}
