package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Store a hashed password in your database
	storedHashedPassword := "$2a$10$n258UHZxExfR3LbAq35xauES/TvGhSiPnAJ77jw6GkLa4Vd4tJHRq"

	// When a user tries to log in, hash the password they entered
	enteredPassword := "gundo"
	enteredHashedPassword, err := bcrypt.GenerateFromPassword([]byte(enteredPassword), bcrypt.DefaultCost
	if err != nil {
		// Handle error
	}

	// Compare the two hashed passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(enteredHashedPassword), Defa\
	if err == nil {
		// Passwords match, proceed with login
		fmt.Println("Login successful!")
	} else {
		// Passwords don't match, deny login attempt
		fmt.Println("Login failed!")
	}
}
