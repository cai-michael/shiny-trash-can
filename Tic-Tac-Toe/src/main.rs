use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {

    // Initialize the board
    println!("Welcome to Tic-Tac-Toe");

    let main_menu_options = "Please choose the game type:\n\
                             1. Play Game vs a friend\n\
                             2. Play Game vs AI (Random)\n\
                             3. Play Game vs AI (MiniMax)\n\
                             4. Quit";
    
    // Loop Starts Here
    loop {
        println!("{}", main_menu_options);

        // Take user options
        // Choose the kind of game
        let mut user_input = String::new();
        
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
            1 => println!("Not implemented yet"),
            2 => println!("Not implemented yet"),
            3 => println!("Not implemented yet"),
            4 => {
                println!("Quitting...");
                break;
            },
            _ => println!("Invalid User Option Selected"),
        }

        // Choose who goes first
        
        // Proceed with game

        //Random AI move
        //let mut ai_move = rand::thread_rng().gen_range(1..9);
        // On a win condition reset to the beginning
    }
    
}