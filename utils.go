package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func getInputPassword() (pwd string, err error) {
	fmt.Print("Enter Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a new line after the password entry

	fmt.Print("Confirm Password: ")
	confirmPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()

	if strings.TrimSpace(string(password)) != strings.TrimSpace(string(confirmPassword)) {
		return "", errors.New("Passwords do not match")
	}

	return strings.TrimSpace(string(password)), nil
}
