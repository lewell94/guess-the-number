package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"text/scanner"
	"strconv"
)

func main() {

	// Generate random int between 1 and 10, using the Unix time in nanoseconds as the seed
	timeNow     := time.Now().UnixNano()
	source      := rand.NewSource(timeNow)
	rand        := rand.New(source)
	magicNumber := rand.Intn(10)

	// Tell the user what they have to do
	fmt.Println("Try and guess the magic number between 0 and 10 (inclusive)...")

	// Create scanner, initialising it to read from Stdin
	var s scanner.Scanner
	s.Init(os.Stdin)

	// Scan the inputted text
	var token rune
	for token != scanner.EOF {

		// Get the rune of the inputted text
		token = s.Scan()

		// Get the inputted text (as string), convert it to an int
		text     := s.TokenText()
		guess, _ := strconv.Atoi(text)

		// Check that the inputted number is between 0 and 10
		if guess > 10 || guess < 0 {
			fmt.Println("Sorry, the magic number is between 0 and 10 (inclusive). Try again!")
		} else {

			// Check to see whether the user has guessed correctly
			if guess == magicNumber {
				fmt.Println("You are correct! Congratulations.")
			} else {
				fmt.Printf("%s %d%s \n", "You are incorrect! The magic number was", magicNumber, ".")
			}
		}

		os.Exit(0)
	}
}