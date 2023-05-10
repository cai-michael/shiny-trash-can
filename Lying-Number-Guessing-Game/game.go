package main

import (
	"fmt"
	"math/rand"
	"strconv"

	Solver "github.com/cai-michael/shiny-trash-can/Lying-Number-Guessing-Game/solver"
)

func main() {
	fmt.Println("Hello! Please select a mode.")
	fmt.Println("1: Normal Guessing Game")
	fmt.Println("2: Compare Strategies in Simulation")
	var menu_choice int = get_menu_input()
	if menu_choice == 1 {
		player_game()
	}
	if menu_choice == 2 {
		compare_strategies_in_simulation(100)
	}
}

func compare_strategies_in_simulation(n int) {
	hundred_approximator_guesses := make([]int, n)
	hundred_approximator_total_guesses := 0

	for i := 0; i < n; i++ {
		// Generate the answer
		var answer int = rand.Intn(100) + 1

		// Run the simulation on different strategies
		hundred_approximator_guesses[i] = Solver.Hundred_approximator_strategy(answer)
		hundred_approximator_total_guesses += hundred_approximator_guesses[i]
	}

	// Print results

	fmt.Printf("Hundred Approximator Strategy Average Guesses:\t%d!", hundred_approximator_total_guesses/n)
}

func player_game() {
	fmt.Println("Hello player! In this game you must guess a number between 1-100.")
	fmt.Println("If your guess is correct you win! But if it isn't I'll tell you if it's higher or lower.")
	fmt.Println("Be careful though, I will lie to you x% of the time where x is the answer.")

	// Generate the answer
	var answer int = rand.Intn(100) + 1
	var guesses int = 0
	var guess int

	for {
		// Take user input
		guess = get_guess_input()
		guesses++

		// Determine if user's guess is correct
		if guess != answer {
			var lie bool = false
			random_chance := rand.Intn(100) + 1
			if random_chance <= answer {
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

func get_guess_input() int {
	for {
		fmt.Println("Enter your guess:")
		var choice int = get_player_input_integer()
		if choice <= 2 && choice >= 1 {
			return choice
		} else {
			fmt.Println("The menu option is not valid.")
		}

	}
}

func get_menu_input() int {
	for {
		fmt.Println("Enter your selection:")
		var guess int = get_player_input_integer()
		if guess <= 100 && guess >= 1 {
			return guess
		} else {
			fmt.Println("The guessed number is not between 1 and 100")
		}

	}
}

func get_player_input_integer() int {
	for {
		var user_input string
		fmt.Scanln(&user_input)
		if input_as_int, err := strconv.Atoi(user_input); err == nil {
			return input_as_int
		} else {
			fmt.Println("The input is not a valid number")
		}
	}

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
