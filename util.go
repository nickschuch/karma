package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Passes the text and looks for a username.
func getUser(t string) (string, error) {
	// Remove all except for characters.
	reg, err := regexp.Compile("[^A-Za-z]+")
	if err != nil {
		log.Fatal(err)
	}
	safe := reg.ReplaceAllString(t, "")
	safe = strings.ToLower(strings.Trim(safe, "-"))
	return safe, nil
}

// Check if the text asked for an increase.
func increaseAmount(t string) int {
	// If the user gets a ++ result.
	if strings.Contains(t, "++") {
		return 1
	}

	// If the user gets a += result.
	if strings.Contains(t, "+=") {
		return findMultiAmount(t)
	}

	return 0
}

// Check if the text asked for a decrease.
func decreaseAmount(t string) int {
	// If the user gets a -- result.
	if strings.Contains(t, "--") {
		return 1
	}

	// If the user gets a -= result.
	if strings.Contains(t, "-=") {
		return findMultiAmount(t)
	}

	return 0
}

// Common handler for "+=" and "-=" strings.
func findMultiAmount(t string) int {
	slice := strings.Split(t, "=")
	// Ensure there is a value.
	if len(slice[1]) > 0 {
		// Ensure we don't have any unwanted characters.
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		replaced := reg.ReplaceAllString(slice[1], "")

		// Convert it to an int for calcuating.
		value, err := strconv.Atoi(replaced)
		if err != nil {
			return 0
		}
		return value
	}

	return 0
}
