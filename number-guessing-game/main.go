package main

import (
	"fmt"
	"math/rand"
)

var chances int
var difficulty int
var guess int

func main() {
	num := rand.Intn(100)
	attempts := 0

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	fmt.Print("Enter your choice:")
	fmt.Scan(&difficulty)

	switch {
	case difficulty == 1:
		chances = 10
		fmt.Println("Great! You have selected the Easy difficulty level.")
	case difficulty == 2:
		chances = 5
		fmt.Println("Great! You have selected the Medium difficulty level.")
	case difficulty == 3:
		chances = 3
		fmt.Println("Great! You have selected the Hard difficulty level.")
	}

	fmt.Println("Let's start the game!")
	//fmt.Printf("number:%d\n", num)

	for i := 0; i <= chances; i++ {
		fmt.Printf("chances: %d\n", chances)
		fmt.Print("Enter your guess:")
		fmt.Scan(&guess)

		if guess > num {
			fmt.Printf("Incorrect! The number is less than %d.\n\n", guess)
			chances -= 1
		} else if guess < num {
			fmt.Printf("Incorrect! The number is greater than %d.\n\n", guess)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n\n", attempts)
			break
		}
		//fmt.Println("i:", i)

		if i == chances {
			fmt.Println("You run out of chances, YOU ARE LOSER.")
		}
		attempts += 1
	}
}
