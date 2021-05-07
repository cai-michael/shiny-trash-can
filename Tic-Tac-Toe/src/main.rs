pub mod board;

use rand::Rng;
use std::cmp::Ordering;
use std::io;
use board::Board;

fn main() {

    // Print welcome message
    println!("Welcome to Tic-Tac-Toe");

    // Initialize the board
    let mut board = Board::new();

    board.change_state(2, X);
    board.print_board();

    let main_menu_options = "Please choose the game type:\n\
                             1. Play Game vs a friend\n\
                             2. Play Game vs AI (Random)\n\
                             3. Play Game vs AI (MiniMax)\n\
                             4. Quit";
    
    let player_order_options = "Who will go first?:\n\
                                1. Player\n\
                                2. AI";
    
    // Loop Starts Here
    loop {
        println!("{}", main_menu_options);

        // Take user options
        // Choose the kind of game
        let mut user_input = String::new();
        let mut ai_enabled = false;
        let mut player1_turn = false;
        
        io::stdin()
            .read_line(&mut user_input)
            .expect("Failed to read line");
        
        let user_option: u32 = match user_input.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Please input a number"); 
                continue;
            },
        };
        
        match user_option {
            1 => {
                println!("Not implemented yet");
                ai_enabled = false;
            },
            2 => {
                println!("Not implemented yet");
            },
            3 => {
                println!("Not implemented yet");
            },
            4 => {
                println!("Quitting...");
                break;
            },
            _ => {
                println!("Invalid User Option Selected");
                continue; // should repeat choice
            },
        }

        // Choose who goes first
        println!("{}", player_order_options);

        io::stdin()
            .read_line(&mut user_input)
            .expect("Failed to read line");

        let user_option: u32 = match user_input.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Please input a number"); 
                continue;
            },
        };

        // Proceed with game


        //Random AI move
        //let mut ai_move = rand::thread_rng().gen_range(1..9);
        // On a win condition reset to the beginning
    }
    
}