/*-------------------------------------------------------------------------
// AUTHOR: Michael Cai
// FILENAME: Board.java
// SPECIFICATION: Simulates the Peg Board Game (Class file for the game board whch keeps track of the pegs)
// FOR: CSE 110- Honors Contract
// TIME SPENT: 4 hours
//-----------------------------------------------------------*/
public class Board {

	private static Hole[] holes;
	private static ValidJumps[] validJumps = new ValidJumps[36];
	private final int BOARD_SIZE = 15; //There are 15 holes in the board
	private int input;
	private final int NOOFVALIDJUMPS = 36; // There are 36 valid jumps

	public Board()
	{
		//Create the board size with the number of holes we want
		holes = new Hole[BOARD_SIZE];

		for (int i = 0; i < BOARD_SIZE; i++)
		{
			holes[i] = new Hole(i);
		}

		initializeJumps();

		//removes the peg from A to have a hole to jump to
		holes[0].removePeg();
	}//end of constructor

	//converts the letter entered into a readable command
	public int convert(String enteredLetter)
	{
		input = (int)(enteredLetter.toLowerCase().charAt(0)-'a');
		return input;
	}//end of method convert

	/*
	//prints the letters corresponding to each input (defunct for test driver)
	public void printPositions()
	{
		System.out.print("\nPegboard Positions\n"
				+ "-----------------------------------\n"
				+ "            A\n"
				+ "           B C\n"
				+ "          D E F\n"
				+ "         G H I J\n"
				+ "        K L M N O\n");

		 //The program will view the letters as numbers due to the conversion
		 //A = 0, B = 2, C = 3...; O = 15

	}//end of method printPositions
	 */
	//prints the current state of the board and the letters corresponding to each jump
	public void printBoard()
	{
		System.out.print("Peg Status     Peg Positions\n"
				+ "    " + holes[0].displayPeg() + "                A\n"
				+ "   " + holes[1].displayPeg() + " " + holes[2].displayPeg() + "              B C\n"
				+ "  " + holes[3].displayPeg() + " " + holes[4].displayPeg() + " " + holes[5].displayPeg() + "            D E F\n"
				+ " " + holes[6].displayPeg() + " " + holes[7].displayPeg() + " " + holes[8].displayPeg() + " " + holes[9].displayPeg() + "          G H I J\n"
				+ holes[10].displayPeg() + " " + holes[11].displayPeg() + " " + holes[12].displayPeg() + " " + holes[13].displayPeg() + " " + holes[14].displayPeg() + "        K L M N O\n");
		/*
		 The program will view the letters as numbers due to the conversion
		 A = 0, B = 2, C = 3...; O = 15
		 */
	}// end of method printBoard

	//determines if the victory condition (one peg remaining) is met
	public boolean victory()
	{
		//loop about the victory condition
		//check the array to see if there is only 1 peg in the board
		int remainingPegs = 0;
		for(int i = 0; i < holes.length; i++)
		{
			if (holes[i].check())
			{
				remainingPegs += 1;
			}
		}
		if (remainingPegs == 1)
		{
			return true;
		}
		else
		{
			return false;
		}
	}//end of method victory
	//Determines if the selected move is valid and returns the result of the move
	public boolean movePeg(int start, int finish)
	{
		boolean validJump = false;
		//cycles through all of the valid jumps to compare to the selected jump
		for (int i = 0; i < NOOFVALIDJUMPS; i ++)
		{
			ValidJumps currentJump = validJumps[i];
			if (start == currentJump.getStart() && finish == currentJump.getEnd())
			{
				//determines if there is a peg to jump over and a hole on the other side for the inital peg to land in
				if (holes[currentJump.getJumped()].check() && !holes[currentJump.getEnd()].check())
				{
					holes[start].removePeg();
					holes[finish].addPeg();
					holes[currentJump.getJumped()].removePeg();
					validJump = true;
				}
			}
		}
		return validJump;
	}
	//resets all of the pegs on the board back to the original positions
	public void reset()
	{
		for (int i = 0; i < BOARD_SIZE; i++)
		{
			holes[i].addPeg();
		}
		holes[0].removePeg();
	}
	//testing method for TestDriver
	public boolean testPeg()
	{
		holes[0].removePeg();
		return holes[0].check();
	}

	//Initializes all of the valid jumps available
	private static void initializeJumps()
	{
		validJumps[0] = new ValidJumps(0, 3, 1);
		validJumps[1] = new ValidJumps(0, 5, 2);
		validJumps[2] = new ValidJumps(1, 6, 3);
		validJumps[3] = new ValidJumps(1, 8, 4);
		validJumps[4] = new ValidJumps(2, 7, 4);
		validJumps[5] = new ValidJumps(2, 9, 5);
		validJumps[6] = new ValidJumps(3, 0, 1);
		validJumps[7] = new ValidJumps(3, 5, 4);
		validJumps[8] = new ValidJumps(3, 12, 7);
		validJumps[9] = new ValidJumps(3, 10, 6);
		validJumps[10] = new ValidJumps(4, 11, 7);
		validJumps[11] = new ValidJumps(4, 13, 8);
		validJumps[12] = new ValidJumps(5, 0, 2);
		validJumps[13] = new ValidJumps(5, 3, 4);
		validJumps[14] = new ValidJumps(5, 12, 8);
		validJumps[15] = new ValidJumps(5, 14, 9);
		validJumps[16] = new ValidJumps(6, 1, 3);
		validJumps[17] = new ValidJumps(6, 8, 7);
		validJumps[18] = new ValidJumps(7, 2, 4);
		validJumps[19] = new ValidJumps(7, 9, 8);
		validJumps[20] = new ValidJumps(8, 1, 4);
		validJumps[21] = new ValidJumps(8, 6, 7);
		validJumps[22] = new ValidJumps(9, 2, 5);
		validJumps[23] = new ValidJumps(9, 7, 8);
		validJumps[24] = new ValidJumps(10, 12, 11);
		validJumps[25] = new ValidJumps(10, 3, 6);
		validJumps[26] = new ValidJumps(11, 4, 7);
		validJumps[27] = new ValidJumps(11, 13, 12);
		validJumps[28] = new ValidJumps(12, 10, 11);
		validJumps[29] = new ValidJumps(12, 14, 13);
		validJumps[30] = new ValidJumps(12, 3, 7);
		validJumps[31] = new ValidJumps(12, 5, 8);
		validJumps[32] = new ValidJumps(13, 11, 12);
		validJumps[33] = new ValidJumps(13, 4, 8);
		validJumps[34] = new ValidJumps(14, 5, 9);
		validJumps[35] = new ValidJumps(14, 12, 13);
	}

}
