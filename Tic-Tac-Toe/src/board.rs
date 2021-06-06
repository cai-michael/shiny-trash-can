use std::io;

pub struct Board {
    dim: usize,
    state: Box<[u8]>,
}

impl Board {
    /// Creates a new 3x3 board.
    ///
    pub fn new() -> Board {
        const DIM: usize = 3;

        Board {
            dim: DIM,
            state: Box::new([0; DIM * DIM]),
        }
    }

    /// Prints the current status of the board
    ///
    pub fn print_board(&self) {
        let mut board = String::new();
        for i in 0..10 {
            if i % 3 == 0 {
                board.push_str(" _____ _____ _____\n");
            } else if i % 3 == 1 {
                board.push_str("|     |     |     |\n");
            } else {
                let first = self.space_to_string(self.state[i - 2], i - 1);
                let second = self.space_to_string(self.state[i - 1], i);
                let third = self.space_to_string(self.state[i], i + 1);
                let line = format!("|  {}  |  {}  |  {}  |\n", first, second, third);
                board.push_str(&line);
            }
        }
        board.push_str("\n");
        println!("{}", board);
    }

    /// Converts board placeholder numbers to their actual symbols
    ///
    /// # Arguments
    ///
    /// * `space` - The number in that space
    /// * `space_num` - The number associated with the space in case no piece is there
    ///
    pub fn space_to_string(&self, space: u8, space_num: usize) -> String {
        match space {
            1 => return "X".to_string(),
            2 => return "O".to_string(),
            _ => return space_num.to_string(),
        }
    }

    /// Changes the state of the board
    ///
    /// # Arguments
    ///
    /// * `position` - The position on the board to be changed
    /// * `piece` - The type of piece to be placed (Blank, X, O)
    ///
    pub fn change_state(&mut self, position: usize, piece: u8) -> bool {
        if position >= self.dim * self.dim {
            return false;
        }

        self.state[position] = piece;
        return true;
    }

    /// Checks if the position is occupied
    ///
    /// # Arguments
    ///
    /// * `position` - The position on the board to be checked
    ///
    pub fn check_space_occupied(&self, position: usize) -> bool {
        if self.state[position] > 0 {
            return true;
        }

        return false;
    }

    /// Gets the input for a move from the user
    ///
    /// # Arguments
    ///
    /// * `player` - The player whos turn it is
    ///
    pub fn get_player_move(&mut self, player: u8) {
        let mut valid_input = false;

        while valid_input == false {
            println!("Which space would you like to play in?");

            let mut user_input = String::new();

            io::stdin()
                .read_line(&mut user_input)
                .expect("Failed to read line");

            let user_option: usize = match user_input.trim().parse() {
                Ok(num) => num,
                Err(_) => {
                    println!("Please input a valid position");
                    continue;
                }
            };

            // Ensure the chosen position is not already occupied
            let chosen_position = user_option - 1;
            if self.check_space_occupied(chosen_position) {
                println!("\nThat space is already occupied.");
                self.print_board();
            } else {
                self.change_state(chosen_position, player);
                valid_input = true;
            }
        }
    }

    /// Determines if any player has won or if there is a tie game
    ///
    /// # Arguments
    ///
    /// * `ai_enabled` - A flag determining if an AI is player 2
    ///
    pub fn check_win_condition(&self, ai_enabled: bool) -> bool {
        let mut winner = 0;

        // Check if there is any three horizontal
        for i in 0..self.dim {
            let j = self.dim * i;
            let mut count = 0;

            for k in j..(j + self.dim) {
                if self.state[j] != 0 && self.state[j] == self.state[k] {
                    count += 1;
                }
            }

            if count == self.dim {
                winner = self.state[j];
            }
        }

        // Check if there is any three vertical
        for i in 0..self.dim {
            let mut count = 0;

            for j in 0..self.dim {
                let k = i + (j * self.dim);
                if self.state[i] != 0 && self.state[i] == self.state[k] {
                    count += 1;
                }
            }

            if count == self.dim {
                winner = self.state[i];
            }
        }

        // Check if there is any three in the forward diagonal
        let mut count = 0;
        for i in 0..self.dim {
            let k = i * (self.dim + 1);

            if self.state[i] != 0 && self.state[i] == self.state[k] {
                count += 1;
            }
        }

        if count == self.dim {
            winner = self.state[0];
        }

        // Check if there is any three in the backward diagonal
        let mut count = 0;
        let starting_position = self.dim - 1;
        for i in 0..self.dim {
            let k = starting_position + (i * (self.dim - 1));

            if self.state[starting_position] != 0 && self.state[starting_position] == self.state[k]
            {
                count += 1;
            }
        }

        if count == self.dim {
            winner = self.state[starting_position];
        }

        // Check if a player has won or if the game is tied
        if winner != 0 {
            self.print_board();
            if ai_enabled && winner == 2 {
                println!("Oops, the AI beat you!");
            } else {
                println!("Congragulations! Player {} has won!", winner);
            }
            return true;
        } else {
            let mut count = 0;
            for i in 0..(self.dim * self.dim) {
                if self.state[i] != 0 {
                    count += 1;
                }
            }
            if count == (self.dim * self.dim) {
                self.print_board();
                println!("Oops, Cat's Game!");
                return true;
            }
        }

        return false;
    }

    /// Resets the board to its original state
    ///
    pub fn reset_board(&mut self) {
        for i in 0..(self.dim * self.dim) {
            self.state[i] = 0
        }
    }
}
