package hangman

import (
	"testing"
)

// testing letterInWord
func testLetterInWord(t *testing.T) {
	word := []string{"b", "o", "b"}
	guess := "b"
	hasLetter := letterInWord(guess, word)
	if !hasLetter {
		t.Errorf("Erro")
	}
}
