package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {

	// Generate random int between 1 and 10, using the Unix time in nanoseconds as the seed
	timeNow     := time.Now().UnixNano()
	source      := rand.NewSource(timeNow)
	rand        := rand.New(source)
	magicNumber := rand.Intn(10)

	// Create a new reader, reading from Stdin
	reader := bufio.NewReader(os.Stdin)

	// Tell the user what they have to do
	fmt.Println("Try and guess the magic number between 1 and 10...")

	// Read the user's input from the reader, convert from int32 rune to int character
	input, _, _ := reader.ReadRune()
	guess       := int(input - '0')

	// Check to see whether the user has guessed correctly
	if guess == magicNumber {
		fmt.Println("You are correct! Congratulations.")
	} else {
		fmt.Printf("%s %d%s \n", "You are incorrect! The magic number was", magicNumber, ".")
	}
}