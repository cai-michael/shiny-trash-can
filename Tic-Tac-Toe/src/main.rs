pub mod board;

//use rand::Rng;
use board::Board;
use std::io;

fn main() {
    // Print welcome message
    println!("Welcome to Tic-Tac-Toe");

    // Initialize the board and game variables
    let mut board = Board::new();
    let mut ai_enabled = false;
    let mut player1_piece = 1;
    let mut player2_piece = 2;

    // Loop Starts Here
    loop {
        // Take user options
        // Choose the kind of game
        let user_option = main_menu();

        match user_option {
            1 => {
                println!("\nPlaying with Two Players!");
                ai_enabled = false;
            }
            2 => {
                println!("\nNot implemented yet");
            }
            3 => {
                println!("\nNot implemented yet");
            }
            4 => {
                println!("\nQuitting...");
                break;
            }
            _ => {
                println!("\nInvalid User Option Selected\n");
                continue;
            }
        }

        let mut player1_turn = determine_order(ai_enabled);
        if !player1_turn {
            player1_piece = 2;
            player2_piece = 1;
        }

        // Proceed with game
        loop {
            board.print_board();

            if player1_turn {
                // Get player 1's input and change the board state
                println!("Player 1 your turn:");
                board.get_player_move(player1_piece);
                player1_turn = false;
            } else {
                if ai_enabled {
                    // Get AI's move
                    //Random AI move
                    //let mut ai_move = rand::thread_rng().gen_range(1..9);
                    println!("The AI is playing...");
                    board.get_player_move(player2_piece);
                } else {
                    // Get player 2's input and change the board state
                    println!("Player 2 your turn:");
                    board.get_player_move(player2_piece);
                }

                player1_turn = true;
            }

            // Check win condition
            if board.check_win_condition(ai_enabled) {
                board.print_board();
                board.reset_board();
                break;
            }
        }

        // On a win condition reset to the beginning
    }
}

/// Prints the main menu and returns the option chosen by the user
///
fn main_menu() -> u32 {
    //let main_menu_options = "Please choose the game type:\n\
    //                         1. Play Game vs a Friend\n\
    //                         2. Play Game vs AI (Random)\n\
    //                         3. Play Game vs AI (MiniMax)\n\
    //                         4. Quit";

    let main_menu_options = "Please choose the game type:\n\
                             1. Play Game vs a Friend\n\
                             4. Quit";

    println!("{}", main_menu_options);

    let mut user_input = String::new();
    let mut user_option: u32 = 0;
    let mut valid_input = false;

    while valid_input == false {
        io::stdin()
            .read_line(&mut user_input)
            .expect("Failed to read line");

        user_option = match user_input.trim().parse() {
            Ok(num) => {
                valid_input = true;
                num
            }
            Err(_) => {
                println!("Please input a number");
                continue;
            }
        };
    }

    return user_option;
}

/// Prints options of who to go first when playing with AI
///
/// # Arguments
///
/// * `ai_enabled` - The flag as to whether or not an AI is enabled
///
fn determine_order(ai_enabled: bool) -> bool {
    let player_order_options = "Who will go first?:\n\
                                1. Player\n\
                                2. AI";

    // Choose who goes first

    let mut user_input = String::new();
    let mut valid_input = false;
    let mut user_goes_first = true;

    if ai_enabled == true {
        println!("{}", player_order_options);

        while valid_input == false {
            io::stdin()
                .read_line(&mut user_input)
                .expect("Failed to read line");

            let user_option = match user_input.trim().parse() {
                Ok(num) => num,
                Err(_) => {
                    println!("Please input a number");
                    continue;
                }
            };

            match user_option {
                1 => {
                    println!("It's your turn!");
                    valid_input = true;
                }
                2 => {
                    println!("The AI is going first...");
                    valid_input = true;
                    user_goes_first = false;
                }
                _ => {
                    println!("Invalid User Option Selected");
                    continue;
                }
            }
        }
    }

    return user_goes_first;
}
