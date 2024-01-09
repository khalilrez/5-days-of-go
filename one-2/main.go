package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	fmt.Println(random.Intn(101))

	fmt.Println("Welcome to the Age Guessing Game!")
	fmt.Println("Think of a number between 1 and 100, and I'll try to guess it.")

	for {
		fmt.Println("Enter 'y' to continue")
		confirmation, err := bufio.NewReader(os.Stdin).ReadString('\n')
		confirmation = strings.TrimSpace(confirmation)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue // Restart loop on error
		}

		// Validate user input
		if confirmation == "y" {
			break // Exit loop if confirmed
		} else {
			return
		}
	}
	// Binary Search to guess age
	low, high := 0, 100
	for low <= high {
		guess := (low + high) / 2.0
		fmt.Println("( current low ", low, "current high ", high, " ) Is it ", guess, " ?")

		fmt.Println("Enter 'h' if higher , 'l' if lower and 'c' if correct :  ")
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading your answer : ", err)
			continue
		}
		answer = strings.TrimSpace(answer)

		switch answer {
		case "h":
			low = guess + 1
		case "l":
			high = guess - 1
		case "c":
			fmt.Println("Yay ! I Got it ")
			return
		default:
			fmt.Println("Insert a valid answer ( 'h' , 'l' or 'c' )")
			continue
		}
	}
}
