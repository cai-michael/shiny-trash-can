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
	guesses := make([]State, 210)
	num_guesses := 0

	num_higher := 0
	for i := 0; i < 10; i++ {
		guesses[i] = get_game_state(100, answer)
		if guesses[i] == Higher {
			num_higher++
		}
		num_guesses++
	}

	current_guess := num_higher * 10
	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			current_guess = current_guess + i
		} else {
			current_guess = current_guess - i
		}

		if current_guess <= 100 && current_guess > 0 {
			game_state := get_game_state(current_guess, answer)
			guesses[i+10] = game_state
			num_guesses++
			if game_state == Correct {
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
	} else if answer > guess {
		if !lie {
			return Higher
		} else {
			return Lower
		}
	}

	return Correct
}
