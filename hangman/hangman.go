package hangman

import "strings"

type Game struct {
	State        string   // game state
	Letters      []string // letters to find
	FoundLetters []string // letters found (Good Guesses)
	UsedLetters  []string // used letter
	TurnLeft     int      //

}

func New(turn int, word string) *Game {
	letter := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letter))
	for i := 0; i < len(letter); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letter,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnLeft:     turn,
	}
	return g
}
