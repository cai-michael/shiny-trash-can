package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Hello player! In this game you must guess a number between 1-100.")
	fmt.Println("If your guess is correct you win! But if it isn't I'll tell you if it's higher or lower.")
	fmt.Println("Be careful though, I will lie to you x% of the time where x is your current guess.")

	// Generate the answer
	var answer int = rand.Intn(100)
	var guesses int = 0
	var guess int

	for {
		print(answer)
		// Take user input
		guess = get_player_input()
		guesses++

		// Determine if user's guess is correct
		if guess != answer {
			var lie bool = false
			random_chance := rand.Intn(100)
			if random_chance <= guess {
				lie = true
			}

			give_player_hint(guess, answer, lie)

		} else {
			// The answer is correct
			break
		}
	}

	fmt.Printf("Congragulations! You guessed the number in %d guesses!", guesses)
}

func get_player_input() int {
	var guess int

	for {
		var user_input string
		fmt.Println("Enter your guess:")
		fmt.Scanln(&user_input)

		if guess_int, err := strconv.Atoi(user_input); err == nil {
			if guess_int <= 100 && guess_int >= 1 {
				guess = guess_int
				break
			} else {
				fmt.Println("The guessed number is not between 1 and 100")
			}
		} else {
			fmt.Println("The input is not a valid number")
		}
	}

	return guess
}

func give_player_hint(guess int, answer int, lie bool) {
	var lower_hint string = "The answer is lower than your guess!"
	var higher_hint string = "The answer is higher than your guess!"

	if answer < guess {
		if !lie {
			fmt.Println(lower_hint)
		} else {
			fmt.Println(higher_hint)
		}
	} else {
		if !lie {
			fmt.Println(higher_hint)
		} else {
			fmt.Println(lower_hint)
		}
	}
}
