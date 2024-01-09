package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	secretNumber := rand.Intn(100) + 1
	attempts := 10
	currentPlayer := "user"
	low, high := 1, 100

	fmt.Println("Welcome to the Number Guessing Game!")

	for attempts > 0 {

		switch currentPlayer {
		case "user":
			fmt.Println("You have ", attempts, " attempts left.")
			var guess int
			fmt.Print("Enter your guess (1-100): ")
			fmt.Scanln(&guess)
			if guess == secretNumber {
				fmt.Println("Congratulations! You guessed the number!")
				return
			} else if guess < secretNumber {
				fmt.Println("Too low!")
			} else {
				fmt.Println("Too high!")
			}
			attempts--
		case "computer":
			guess := rand.Intn(high-low+1) + low
			fmt.Printf("Computer's guess: %d\n", guess)
			if guess == secretNumber {
				fmt.Println("The computer guessed the number!")
				return
			} else if guess < secretNumber {
				fmt.Println("Computer's guess was too low.")
				low = guess + 1
			} else {
				fmt.Println("Computer's guess was too high.")
				high = guess - 1
			}
		}
		println("\n\n")

		currentPlayer = switchPlayer(currentPlayer)
	}

	fmt.Println("Game over! No one guessed the number.")
	fmt.Printf("The secret number was %d.\n", secretNumber)
}

func switchPlayer(player string) string {
	if player == "user" {
		return "computer"
	}
	return "user"
}
