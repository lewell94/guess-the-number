package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
	"bufio"
)

func generateRandom() int {

	// Generate random int between 0 and 10 inclusive, using the Unix time in nanoseconds as the seed
	timeNow     := time.Now().UnixNano()
	source      := rand.NewSource(timeNow)
	rand        := rand.New(source)
	magicNumber := rand.Intn(10)

	return magicNumber
}

func checkInput(input string, magicNumber int) int {

	// Convert input string to an int
	guess, _ := strconv.Atoi(input)

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

	return 0
}

func main() {

	// Get the magic number
	magicNumber    := generateRandom()

	// Create a new reader, reading from stdin
	reader         := bufio.NewReader(os.Stdin)

	// Tell the user what they have to do
	fmt.Println("Try and guess the magic number between 1 and 10...")

	// Read the inputted string, stop reading at a newline
	inputString, _ := reader.ReadString(10)

	// Get length of inputted string
	inputLength    := len(inputString)

	// Check to see if input is valid - that it contains only numbers, no letters or punctuation
	valid          := true

	// Loop through the inputted string, it there are any letters
	// or punctuation set valid to false. Ignore the last byte by
	// stopping the loop at inputLength-1, as this will always be 
	// 10, for the newline
	for i := 0; i < inputLength-1; i++ {

		// TO-DO: Check for a hyphen, as this would be used for negative 
		// numbers, therefore needs to show "magic number is between 0
		// and 10" message instead of "no letters or punctuation" message

		if inputString[i] < 48 || inputString[i] > 57 {
			valid = false
		}
	}

	// If the user has inputted any non-numbers, tell them to try again
	if !valid {
		fmt.Println("Sorry, the magic number is a number, and therefore does not contain any letters or punctuation! Try again!")
	} else {

		// Create an empty string, which will eventually be passed into the checkInput function
		var input string

		// We need to pass in the inputted string to the checkInput function, but without the
		// newline at the end, therefore, loop through it up until inputLength-1, and append
		// the characters to the above empty string
		for i := 0; i < inputLength-1; i++ {

			// Get the int from the byte of each character in the string
			// (byte is given to us by inputString[i])
			charAsInt := int(inputString[i] - '0')

			// Convert the int to a string for appending the input
			chasAsStr := strconv.Itoa(charAsInt)

			input += chasAsStr
		}
		
		// Check the input against the magic number
		checkInput(input, magicNumber)
	}
}