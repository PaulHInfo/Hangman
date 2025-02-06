package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) //Keyboard reading

func ReadGuess() (guess string, err error) {
	ok := false

	for !ok {
		fmt.Println("Choose a letter : ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Println("Invalide letter")
			continue
		}
		ok = true
	}

	return guess, nil
}
