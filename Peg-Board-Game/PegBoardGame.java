/*-------------------------------------------------------------------------
// AUTHOR: Michael Cai
// FILENAME: PegBoardGame.java
// SPECIFICATION: Simulates the Peg Board Game
// FOR: CSE 110- Honors Contract
// TIME SPENT: 11 hours (1 for main, 4 conceptually, 6 on classes)
//-----------------------------------------------------------*/
import java.util.*;

public class PegBoardGame {

	public static void main(String[] args)
	{

		int selectedStartPeg;
		int selectedEndingPeg;
		Scanner scan = new Scanner(System.in);
		Board board = new Board();
		int selection;
		//create the pegs
		System.out.println("Jump over each peg with another peg to remove it\n"
				+ "Try to leave as few pegs as possible\n"
				+ "An X means a peg is there, an O means the hole is empty\n");
		board.printBoard();
		System.out.println("\nUse the corresponding letters to pick the pegs and holes\n");

		//loop until victory condition
		while (!board.victory())
		{
			do 
			{
				System.out.println("\nPress 1 to move a peg\n"
						+ "Press 9 to reset the board");
				selection = scan.nextInt();
				if (selection != 1 && selection != 9)
				{
					System.out.println("Error: Invalid Input");
				}
			} while (selection != 1 && selection != 9);
			switch (selection)
			{
			case 1:
				System.out.println("Which peg do you want to move?");
				selectedStartPeg = board.convert(scan.next());
				System.out.println("Where do you want to move the peg?");
				selectedEndingPeg = board.convert(scan.next());
				if (board.movePeg(selectedStartPeg, selectedEndingPeg))
					board.printBoard();
				else
				{
					System.out.println("Error: Invalid Jump");
					board.printBoard();
				}
				break;
			case 9:
				board.reset();
				board.printBoard();
				break;
			}
		}
		System.out.println("Congratulations! You Won!");
		scan.close();//close the scanner
	}//end of main method
}