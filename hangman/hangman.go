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
		State:        " ",
		Letters:      letter,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnLeft:     turn,
	}
	return g
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = "AG"
	} else if letterInWord(guess, g.Letters) {
		g.State = "GG"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "W"
		}
	} else {
		g.State = "BG"
		g.TurnLeft = g.TurnLeft - 1
		g.UsedLetters = append(g.UsedLetters, guess)
		if g.TurnLeft <= 0 {
			g.State = "L"
		}

	}

}
func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}

}
func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}
