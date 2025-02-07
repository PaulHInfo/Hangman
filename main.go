package main

import (
	"fmt"
	"hangman/dictionnary"
	"hangman/hangman"
	"os"
)

func main() {

	err := dictionnary.Load("data/dico.txt")

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	w := dictionnary.PickWord()
	fmt.Println()
	g := hangman.New(8, w)
	hangman.DrawWelcom()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "W", "l":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Println("ERROR")
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
