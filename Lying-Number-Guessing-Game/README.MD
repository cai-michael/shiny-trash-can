# Lying-Number-Guessing-Game

The game generates a number between 0 and 100 inclusive. 

The player guesses a number

The game will respond by either saying higher or lower or correct. BUT if the player incorrectly guesses then the game has a percent chance to lie and say the opposite of the truth. This percent chance is determined by the number being guessed. Ie if the number is 72, then there is a 72% chance of a lie.

P(higher | number is higher) = 1 - number/100
P(lower | number is lower) = 1 - number/100