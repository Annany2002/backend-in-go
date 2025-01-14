package utils

import (
	"log"
)

// CheckError is a helper function to check for errors and log them
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// LogInfo is a helper function to log informational messages
func LogInfo(message string) {
	log.Println("INFO: " + message)
}

// LogWarning is a helper function to log warning messages
func LogWarning(message string) {
	log.Println("WARNING: " + message)
}

// LogError is a helper function to log error messages
func LogError(message string) {
	log.Println("ERROR: " + message)
}
