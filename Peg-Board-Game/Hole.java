/*-------------------------------------------------------------------------
// AUTHOR: Michael Cai
// FILENAME: Hole.java
// SPECIFICATION: Simulates the Peg Board Game (Class file for the individual pegs)
// FOR: CSE 110- Honors Contract
// TIME SPENT: 1 hour
//-----------------------------------------------------------*/

public class Hole {
	
	private int holeNo; //The corresponding position of the hole
	private boolean peg = true; //Whether or not a peg is in that position
	
	public Hole(int number)
	{
		holeNo = number;
		if (holeNo == 0)
				{
					peg = false;
				}
	}
	
	//Method used to return information on whether or not a peg is in that position
	public boolean check()
	{
		if (peg == true) //Peg is in that position
		{
		return true;
		}
		else //Peg is not in that position
		{
		return false;
		}
	}//end of method "check"
	//removes a peg from the hole
	public void removePeg()
	{
		peg = false;
	}//end of method removePeg
	//adds a peg to the hole
	public void addPeg()
	{
		peg = true;
	}//end of method addPeg
	//method used to convey information on if a peg is in a hole or not
	public String displayPeg()
	{
		if (this.check() == true)
		{
			return "X";
		}
		else
		{
			return "O";
		}
	}
}

