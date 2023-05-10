package solver

import (
	"math/rand"
)

type State int

const (
	Higher  State = 0
	Lower   State = 1
	Correct State = 2
)

func Hundred_approximator_strategy(answer int) int {
	guesses := make([]State, 10)
	num_guesses := 0

	num_higher := 0
	for i := 0; i < 10; i++ {
		guesses[i] = get_game_state(100, answer)
		if guesses[i] == Higher {
			num_higher++
		}
		num_guesses++
	}

	starting_point := num_higher * 10
	for i := 0; i < 200; i++ {

		var current_guess int
		if i%2 == 0 {
			current_guess = starting_point + i
		} else {
			current_guess = starting_point - i
		}

		if current_guess <= 100 && current_guess > 0 {
			guesses[i] = get_game_state(current_guess, answer)
			num_guesses++
			if guesses[i] == Correct {
				break
			}
		}
	}

	return num_guesses
}

func get_game_state(guess int, answer int) State {
	var lie bool = false
	random_chance := rand.Intn(100) + 1
	if random_chance <= answer {
		lie = true
	}

	if answer < guess {
		if !lie {
			return Lower
		} else {
			return Higher
		}
	} else {
		if !lie {
			return Higher
		} else {
			return Lower
		}
	}
}
