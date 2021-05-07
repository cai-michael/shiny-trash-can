pub struct Board {
    width: usize,
    height: usize,
    state: Box<[u8]>,
}

impl Board {
    pub fn new() -> Board {
        const Width: usize = 3;
        const Height: usize = 3;

        Board {
            width: Width,
            height: Height,
            state: Box::new([0;Width*Height]),
        }
    }

    /// Prints the current status of the board
    ///
    pub fn print_board(&self) {
        let mut board = String::new();
        for i in (0..10) {
            if i % 3 == 0 { 
                board.push_str(" _____ _____ _____\n");
            }
            else if i % 3 == 1 {
                board.push_str("|     |     |     |\n");
            }
            else {
                let first = self.space_to_string(self.state[i - 2]);
                let second = self.space_to_string(self.state[i - 1]);
                let third = self.space_to_string(self.state[i]);
                let line = format!("|  {}  |  {}  |  {}  |\n", first, second, third);
                board.push_str(&line);
            }
        }
        board.push_str("\n");
        println!("{}", board);
    }

    pub fn space_to_string(&self, space: u8) -> String {
        match space {
            _ => return " ".to_string(),
            1 => return "X".to_string(),
            2 => return "O".to_string(),
        }
    }

    /// Changes the state of the board
    ///
    /// # Arguments
    ///
    /// * `position` - The position on the board to be changed
    /// * `piece` - The type of piece to be placed (Blank, X, O)
    ///
    pub fn change_state(&mut self, position: usize, piece: Space) -> bool {
        if position < self.width * self.height {
            return false;
        }
        
        self.state[position] = 1;
        return true
    }
}