// Assignment #: Honors Contract
//         Name: Michael Cai
//    StudentID: 1212683899
//      Lecture: 8:35 AM MWF
//  Description: Initializes the applet

import javax.swing.*;
import java.util.*;
import java.awt.*;

public class ConnectFour extends JApplet{
	
	private int APPLET_WIDTH = 700, APPLET_HEIGHT = 700;
	private Board board;
	
	public void init()
	{
		board = new Board();
		getContentPane().add(board);
		setSize(APPLET_WIDTH, APPLET_HEIGHT);
	}
}
