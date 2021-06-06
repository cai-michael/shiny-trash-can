% Format a(Row, Column, Number)
% Each number 1..9 is assigned to one cell in each box
1 {a(X,Y,N):X=1..9,Y=1..9,X1<=X,X<=X1+2,Y1<=Y,Y<=Y1+2} 1
:- N=1..9, X1 = 3*(0..2)+1, Y1 = 3*(0..2)+1.

% No two different columns given a row and a number
:- a(X,Y,N), a(X,Y1,N), Y!=Y1.

% No two different rows given a column and a number
:- a(X,Y,N), a(X1,Y,N), X!=X1.

% No two different numbers given a row and a column
:- a(X,Y,N), a(X,Y,N1), N!=N1.

% Greater-Than Variant
% Some positions must have numbers greater than other positions
:- a(R,C,N), a(R1,C1,N1), gt(R,C,R1,C1), N <= N1.

% Greater Than Sudoku
gt(1,2,1,1).
 
gt(1,3,1,2).
gt(1,3,2,3).

gt(1,4,1,5).
 
gt(1,6,1,5).
gt(1,6,2,6).

gt(1,7,2,7).
 
gt(1,8,1,7).
gt(1,8,2,8).
 
gt(1,9,1,8).
gt(1,9,2,9).

gt(2,1,1,1).

gt(2,2,2,1). 
gt(2,2,2,3).
gt(2,2,1,2).
gt(2,2,3,2).

gt(2,3,3,3).

gt(2,4,1,4).
gt(2,4,3,4).
 
gt(2,5,2,4). 
gt(2,5,2,6).
gt(2,5,1,5).
gt(2,5,3,5).

gt(2,6,3,6).
 
gt(2,8,2,7).
 
gt(2,9,2,8).
gt(2,9,3,9).

gt(3,1,3,2).
gt(3,1,2,1).
 
gt(3,3,3,2).
 
gt(3,4,3,5).
 
gt(3,5,3,6).

gt(3,7,2,7). 
gt(3,7,3,8).

gt(3,8,2,8).
 
gt(3,9,3,8).

gt(4,1,4,2).
gt(4,1,5,1).
 
gt(4,3,4,2).
gt(4,3,5,3).
 
gt(4,5,4,4).
 
gt(4,6,4,5).
gt(4,6,5,6).
 
gt(4,7,4,8).
 
gt(4,9,4,8).

gt(5,2,5,1). 
gt(5,2,5,3).
gt(5,2,4,2).
gt(5,2,6,2).
 
gt(5,4,5,5).
gt(5,4,4,4).
gt(5,4,6,4).
 
gt(5,5,4,5).
gt(5,5,6,5).

gt(5,6,5,5).

gt(5,7,4,7). 
gt(5,7,5,8).
 
gt(5,8,5,9).
gt(5,8,4,8).

gt(5,9,4,9).

gt(6,1,5,1).

gt(6,2,6,1). 
gt(6,2,6,3).

gt(6,3,5,3).
 
gt(6,5,6,4).
 
gt(6,6,6,5).
gt(6,6,5,6).

gt(6,7,5,7).
 
gt(6,8,6,7). 
gt(6,8,6,9).
gt(6,8,5,8).

gt(6,9,5,9).

gt(7,1,7,2).
gt(7,1,8,1).
 
gt(7,3,7,2).
gt(7,3,8,3).
 
gt(7,4,7,5).
gt(7,4,8,4).
 
gt(7,6,7,5).
gt(7,6,8,6).

gt(7,7,8,7).
 
gt(7,8,7,7). 
gt(7,8,7,9).

gt(8,1,8,2).
gt(8,1,9,1).
 
gt(8,2,8,3).
gt(8,2,7,2).

gt(8,5,7,5). 
gt(8,5,8,4). 
gt(8,5,8,6).

gt(8,6,9,6).

gt(8,7,9,7).

gt(8,8,7,8). 
gt(8,8,8,7).
gt(8,8,9,8).
 
gt(8,9,8,8).
gt(8,9,7,9).
gt(8,9,9,9).

gt(9,2,9,1). 
gt(9,2,9,3).
gt(9,2,8,2).

gt(9,3,8,3).

gt(9,4,8,4).
 
gt(9,5,9,4). 
gt(9,5,9,6).
gt(9,5,8,5).
 
gt(9,8,9,7).
 
gt(9,9,9,8).