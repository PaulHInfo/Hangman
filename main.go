package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {

	hangman.DrawWelcom()

	g := hangman.New(8, "Simon")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}
	fmt.Println(l)

}
