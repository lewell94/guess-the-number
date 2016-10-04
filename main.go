package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"text/scanner"
	"strconv"
	"regexp"
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

		// Get the inputted text (as string)
		text := s.TokenText()
		
		// If the user inputs the letter q, quit
		if text == "q" {
			os.Exit(0)
		}

		// Check to see if the input only contains numbers
		onlyNumbers, _ := regexp.MatchString("^[0-9]*$", text)

		if onlyNumbers {

			// Convert input string to an int
			guess, _ := strconv.Atoi(text)

			// Check that the inputted number is between 0 and 10
			if guess > 10 || guess < 0 {
				fmt.Println("Sorry, the magic number is between 0 and 10 (inclusive). Try again!")
			} else {

				// Check to see whether the user has guessed correctly
				if guess == magicNumber {
					fmt.Println("You are correct! Congratulations!")
					os.Exit(0)
				} else {
					fmt.Printf("You are incorrect! Feel free to try again, or enter 'q' at any time give up!\n")
				}
			}
		} else {
			fmt.Println("Sorry, the magic number is a number, and therefore does not contain any letters or punctuation! Try again!")
		}
	}
}