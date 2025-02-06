package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {
	g := hangman.New(8, "Sison")
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
