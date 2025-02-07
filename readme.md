# HANGMAN
Hangman game in Go. The program randomly selects a word from the dictionary (```data/dico.txt```). Feel free to modify this word list.

## Rules

The player must guess all the letters of the word. They have eight attempts. If they fail, they will be hanged.

		  _________
		  |        |
		  |        0
		  |       /|\
		  |        |
		  |       /|\
		  |		
		__|___
		|     |_____________
		|	           |		
		|	           |
		|__________________|

## Commandes
first go to the folder of the game and :
```go build main.go```
```go run main.go```