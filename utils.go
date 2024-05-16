package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func getInputPassword() (pwd string) {
	fmt.Print("Enter Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading password")
		return ""
	}
	fmt.Println() // Print a new line after the password entry

	fmt.Print("Confirm Password: ")
	confirmPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading password")
		return ""
	}
	fmt.Println()

	if strings.TrimSpace(string(password)) != strings.TrimSpace(string(confirmPassword)) {
		fmt.Println("Passwords do not match.")
		return ""
	}

	return strings.TrimSpace(string(password))
}
