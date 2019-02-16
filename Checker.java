// Assignment #: 9
//         Name: Michael Cai
//    StudentID: 1212683899
//      Lecture: 8:35 AM MWF
//  Description: Contains all of the attributes of a checker

public class Checker {

	//0 - no color
	//1 - red
	//2 - yellow
	private int color;
	private boolean placed;
	
	//Constructor
	
	public Checker()
	{
		placed = false;
		color = 0;
	}
	
	/*********************
	 * Accessor Methods *
	 *******************/
	public int getColor()
	{
		return color;
	}
	public boolean getPlaced()
	{
		return placed;
	}
	
	//Changes the checker color to the current turn's color and "places" it
	public void placeChecker(int placedColor)
	{
		color = placedColor;
		placed = true;
	}
	//unplaces the checker by making it invisible in the array
	public void removeChecker()
	{
		color = 0;
		placed = false;
	}
	
}
