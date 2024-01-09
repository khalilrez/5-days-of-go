package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get name
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	// Get gender (more inclusive options)
	fmt.Print("What is your gender identity? (man/woman/prefer not to say) : ")
	gender, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading gender:", err)
		return
	}

	// Get age with validation
	fmt.Print("How old are you? (Including this year) : ")
	age, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading age:", err)
		return
	}
	age = strings.TrimSpace(age)
	intAge, errAge := strconv.Atoi(age)
	if errAge != nil {
		fmt.Println("Invalid age entered. Please enter a number.", errAge)
		return
	}

	// Calculate birth year
	now := time.Now()
	birthYear := now.Year() - intAge

	// Construct greeting with inclusive language
	var greeting string
	switch strings.TrimSpace(gender) {
	case "man":
		greeting = "Hello Sir!"
	case "woman":
		greeting = "Hello Madam!"
	default:
		greeting = "Hello!"
	}

	fmt.Println(greeting + "\nYour name is " + strings.TrimSpace(name) + ", and you were born in: " + strconv.Itoa(birthYear))
}
