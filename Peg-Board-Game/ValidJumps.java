/*-------------------------------------------------------------------------
// AUTHOR: Michael Cai
// FILENAME: Hole.java
// SPECIFICATION: Simulates the Peg Board Game (Format to keep the list of all valid and possible jumps)
// FOR: CSE 110- Honors Contract
// TIME SPENT: 20 minutes
//-----------------------------------------------------------*/
public class ValidJumps {

	private int startPeg;
	private int endPeg;
	private int jumpedPeg;
	public ValidJumps(int start, int end, int middle)
	{
		startPeg = start;
		endPeg = end;
		jumpedPeg = middle;
	}
	//returns the peg where the jump is starting from
	public int getStart()
	{
		return startPeg;
	}
	//returns the peg where the jump is going to
	public int getEnd()
	{
		return endPeg;
	}
	//returns the peg that is being jumped over
	public int getJumped()
	{
		return jumpedPeg;
	}
}
