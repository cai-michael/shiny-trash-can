IDENTIFICATION DIVISION.
PROGRAM-ID. FINANCIAL-ACCOUNTING-I.

DATA DIVISION.     
WORKING-STORAGE SECTION.
01 OPTION     PIC 9(2).
01 REVENUE    PIC S9(9).
01 EXPENSES   PIC S9(9).
01 NETVALUE   PIC S9(9).

PROCEDURE DIVISION.
       PERFORM DISPLAYMENU THRU SELECTION UNTIL OPTION = 99
       STOP RUN.

DISPLAYMENU.
*> Shows the currently available options
       DISPLAY "Welcome! Pick an ACC232 option".
       DISPLAY "1. Net Value Calculator".

SELECTION.
*> Makes the menu selection
       ACCEPT OPTION.
       IF OPTION = 1
              PERFORM FINDNETVALUE.
       ENDIF.

FINDNETVALUE.
*> Calculates Net Value
       DISPLAY "Enter Revenue".
              ACCEPT REVENUE.
              DISPLAY "Enter Expenses".
              ACCEPT EXPENSES.
              SUBTRACT EXPENSES FROM REVENUE GIVING NETVALUE.
              IF EXPENSES > REVENUE THEN
                     DISPLAY "Net Loss: "NETVALUE
              ELSE
                     DISPLAY "Net Income: "NETVALUE
              END-IF.

END PROGRAM FINANCIAL-ACCOUNTING-I.