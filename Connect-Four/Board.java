// Assignment #: Honors Contract
//         Name: Michael Cai
//    StudentID: 1212683899
//      Lecture: 8:35 AM MWF
//  Description: Draws the GUI and contains the playing logic
import java.awt.*;
import java.awt.event.*;
import javax.swing.*;

public class Board extends JPanel{

	private Checker[][] checkers;
	private int turnColor, lastPlacedColumn, lastPlacedRow;
	private final int ROWS = 6;
	private final int COLUMNS = 7;
	/*
	 CHECKERWIDTH and CHECKERHEIGHT control the size of the checkers while SPACE_FROM_THE_TOP
	 controls the amount of space between the top of the board and the first checker placed 
	 */
	private final int SPACE_FROM_THE_TOP = 520, CHECKERWIDTH = 95, CHECKERHEIGHT = 95;
	private boolean victory = false, ai = false, aiTurn = false;
	private JButton column1, column2, column3, column4, column5, column6, column7, clear, undo, playAI; 
	private JLabel notifications;

	//Constructor
	public Board()
	{
		//initialize the array and start the game
		checkers = new Checker[COLUMNS][ROWS];
		for (int i = 0; i < COLUMNS; i++)
		{
			for (int j = 0; j < ROWS; j++)
			{
				checkers[i][j] = new Checker();
			}
		}
		turnColor = 1;

		//creates the GUI
		initializeGUI();
	}

	//GUI setup
	public void initializeGUI()
	{
		//Creates the panels and labels
		JPanel buttons = new JPanel();
		JPanel options = new JPanel();
		JPanel columnSelector = new JPanel();
		JPanel notificationBar = new JPanel();
		notifications = new JLabel();

		//sets layouts
		buttons.setLayout(new GridLayout(2,1));
		options.setLayout(new GridLayout(1,3));
		columnSelector.setLayout(new GridLayout(1,COLUMNS));

		//adds the buttons
		column1 = new JButton("1");
		column2 = new JButton("2");
		column3 = new JButton("3");
		column4 = new JButton("4");
		column5 = new JButton("5");
		column6 = new JButton("6");
		column7 = new JButton("7");
		columnSelector.add(column1);
		columnSelector.add(column2);
		columnSelector.add(column3);
		columnSelector.add(column4);
		columnSelector.add(column5);
		columnSelector.add(column6);
		columnSelector.add(column7);
		clear = new JButton("Clear");
		undo = new JButton("Undo");
		playAI = new JButton ("Play vs AI");
		options.add(clear);
		options.add(undo);
		options.add(playAI);
		notificationBar.add(notifications);

		//adds their respective listeners
		ButtonListener listener = new ButtonListener();
		column1.addActionListener(listener);
		column2.addActionListener(listener);
		column3.addActionListener(listener);
		column4.addActionListener(listener);
		column5.addActionListener(listener);
		column6.addActionListener(listener);
		column7.addActionListener(listener);
		clear.addActionListener(listener);
		undo.addActionListener(listener);
		playAI.addActionListener(listener);

		//adds the elements the the applet
		this.setLayout(new BorderLayout());
		notificationBar.setLayout(new FlowLayout());
		buttons.add(columnSelector);
		buttons.add(options);
		this.add(buttons, BorderLayout.SOUTH);
		this.add(notificationBar, BorderLayout.CENTER);
	}

	/********************
	 * Accessor Methods *
	 *********************/
	public boolean getVictory()
	{
		return victory;
	}

	public int getTurnColor()
	{
		return turnColor;
	}

	//changes if the AI is on or off
	private void toggleAI()
	{
		if (ai == false)
		{
			ai = true;
			aiTurn = true;
			notifications.setText("Ai turned on");
		}
		else
		{
			ai = false;
			aiTurn = false;
			notifications.setText("Ai turned off");
		}
	}

	//checks to see where the next available row to place the checker is in a row
	//otherwise returns a prompt saying the move was invalid 
	public void placeChecker(int column, int placedColor)
	{
		boolean validPlacement = false;
		int nextOpenSlot = 0;
		for (int i = ROWS - 1; i >= 0; i--)
		{
			if (checkers[column][i].getPlaced() == false)
			{
				validPlacement = true;
				nextOpenSlot = i;
			}
		}
		if (validPlacement == true)
		{
			checkers[column][nextOpenSlot].placeChecker(placedColor);
			victory = checkVictory(column, nextOpenSlot, placedColor);
			lastPlacedColumn = column;
			lastPlacedRow = nextOpenSlot;
			if (victory == true)
			{
				victory(placedColor);
			}
			else if (placedColor == 1)
			{
				turnColor = 2;
				notifications.setText("");
			}
			else
			{
				turnColor = 1;
				notifications.setText("");
			}
		}
		else
		{
			notifications.setText("Invalid Move");
		}
		if (victory == false)
		{
			if (ai == true && aiTurn == true)
			{
				aiTurn = false;
				do 
				{
					placeChecker((int)(Math.random() * (6 + 1)), turnColor);
				} while (notifications.getText().equals("Invalid Move"));
			}
			else if (ai == true)
			{
				aiTurn = true;
			}
		}
		repaint();
	}

	//clears the board of the checkers
	private void clear()
	{
		//clears the board of all checkers and wipes the page
		for (int i = 0; i < COLUMNS; i++)
		{
			for (int j = 0; j < ROWS; j++)
			{
				checkers[i][j].removeChecker();
			}
		}
		turnColor = 1;
		victory = false;
		notifications.setText("Board Cleared");
		repaint();
	}

	//undoes up to the last move done
	private void undo(int lastColumn, int lastRow)
	{
		//changes the turn back to the color before
		if (checkers[lastPlacedColumn][lastPlacedRow].getColor() != 0)
		{
			turnColor = checkers[lastPlacedColumn][lastPlacedRow].getColor();
		}

		//determines where the last column and row placed were and erases that checker
		if (checkers[lastPlacedColumn][lastPlacedRow].getColor() != 0)
		{
			checkers[lastPlacedColumn][lastPlacedRow].removeChecker();
		}
		victory = false;
		notifications.setText("Last Move Undone");
		repaint();
	}


	//displays the victory message for the appropriate color
	private void victory(int color)
	{
		if (color == 1)
			notifications.setText("Congragulations Red! You Won!");
		else 
			notifications.setText("Congragulations Yellow! You Won!");
	}

	//runs through the row and column the placed checker was in to determine if 4 in a row was found
	private boolean checkVictory(int column, int row, int color)
	{
		int fourInARow = 0;
		//checks for horizontal victory
		try
		{
			for (int i = 0; i <= COLUMNS; i++)
			{
				if (checkers[i][row].getColor() == color && fourInARow != 4)
				{
					fourInARow++;
				}
				else if(fourInARow != 4)
				{
					fourInARow = 0;
				}
			}
		}
		catch (ArrayIndexOutOfBoundsException e)
		{

		}
		//checks for vertical victory
		try 
		{
			if (fourInARow != 4)
			{
				fourInARow = 0;
			}
			for (int i = 0; i <= ROWS; i++)
			{
				if (checkers[column][i].getColor() == color && fourInARow != 4)
				{
					fourInARow++;
				}
				else if(fourInARow != 4)
				{
					fourInARow = 0;
				}
			}
		}
		catch (ArrayIndexOutOfBoundsException e)
		{

		}
		//checks for diagonal victory from bottom left to top right
		try
		{
			int startRow = row;
			int startColumn = column;
			//moves the starting slot to the most bottom left checker from where the last checker was placed
			while (startColumn != 0 && startRow != 0)
			{
				startColumn--;
				startRow--;
			}
			//moves up and right one checker at a time
			for (int i = startColumn; i < COLUMNS; i++)
			{
				if (checkers[i][startRow].getColor() == color && fourInARow != 4)
				{
					fourInARow++;
				}
				else if(fourInARow != 4)
				{
					fourInARow = 0;
				}
				if (startRow < ROWS)
					startRow++;
				else
					break;
			}
		}
		catch (ArrayIndexOutOfBoundsException e)
		{

		}
		//checks for diagonal victory from top left to bottom right
		try
		{
			int startRow = row;
			int startColumn = column;
			while (startColumn != 0 && startRow != ROWS)
			{
				startColumn--;
				startRow++;
			}
			for (int i = startColumn; i < COLUMNS; i++)
			{
				if (checkers[i][startRow].getColor() == color && fourInARow != 4)
				{
					fourInARow++;
				}
				else if(fourInARow != 4)
				{
					fourInARow = 0;
				}
				if (startRow > 0)
					startRow--;
				else
					break;
			}
		}
		catch (ArrayIndexOutOfBoundsException e)
		{

		}

		//returns whether the victory value is true
		if (fourInARow == 4)
		{
			return true;
		}
		else
		{
			return false;
		}
	}

	//Button listener to determine the button and add the checker 
	private class ButtonListener implements ActionListener
	{
		public void actionPerformed(ActionEvent event)
		{
			if (victory == false)
			{
				if (event.getSource() == column1)
				{
					placeChecker(0, turnColor);
				}
				if (event.getSource() == column2)
				{
					placeChecker(1, turnColor);
				}
				if (event.getSource() == column3)
				{
					placeChecker(2, turnColor);
				}
				if (event.getSource() == column4)
				{
					placeChecker(3, turnColor);
				}
				if (event.getSource() == column5)
				{
					placeChecker(4, turnColor);
				}
				if (event.getSource() == column6)
				{
					placeChecker(5, turnColor);
				}
				if (event.getSource() == column7)
				{
					placeChecker(6, turnColor);
				}
			}
			if (event.getSource() == clear)
			{
				clear();
			}
			if (event.getSource() == playAI)
			{
				toggleAI();
			}
			if (event.getSource() == undo)
			{
				undo(lastPlacedColumn, lastPlacedRow);
			}
		}
	}

	public void paint(Graphics Page)
	{
		super.paint(Page);
		for (int i = 0; i < COLUMNS; i++)
		{
			for (int j = 0; j < ROWS; j++)
			{
				int spaceFromRight = i * 100;
				int spaceUp = j * 100;
				int verticalPosition = SPACE_FROM_THE_TOP - spaceUp;
				Page.setColor(Color.white);
				Page.fillOval(spaceFromRight, verticalPosition, CHECKERWIDTH, CHECKERHEIGHT);
			}
		}
		for (int i = 0; i < COLUMNS; i++)
		{
			for (int j = 0; j < ROWS; j++)
			{
				if (checkers[i][j].getColor() == 0)
					break;
				else
				{
					int spaceFromRight = i * 100;
					int spaceUp = j * 100;
					int verticalPosition = SPACE_FROM_THE_TOP - spaceUp;
					if (checkers[i][j].getColor() == 1)
						Page.setColor(Color.RED);
					else
						Page.setColor(Color.YELLOW);
					Page.fillOval(spaceFromRight, verticalPosition, CHECKERWIDTH, CHECKERHEIGHT);
				}
			}
		}
	}
}
