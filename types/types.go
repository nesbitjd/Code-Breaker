package types

import (
	"strings"

	"gorm.io/gorm"
)

type HangmanDB struct {
	gorm.Model
	Word     string
	Failures int
	Guesses  string
}

type Hangman struct {
	Word     string
	Failures int
	Guesses  Guesses
}

type Guesses []string

// ToString converts Guesses into a string
func (g *Guesses) ToString() string {

	return strings.Join(*g, ", ")
}

// GuessesFromString converts a string into Guesses
func GuessesFromString(g string) Guesses {

	return strings.Split(g, ", ")
}

// HangmanToDB converts a Hangman to a HangmanDB
func (h *Hangman) HangmanToDB() HangmanDB {
	return HangmanDB{
		Word:     h.Word,
		Failures: h.Failures,
		Guesses:  h.Guesses.ToString(),
	}
}

// DBtoHangman converts a HangmanDB to a Hangman
func (h *HangmanDB) DBtoHangman() Hangman {
	return Hangman{
		Word:     h.Word,
		Failures: h.Failures,
		Guesses:  GuessesFromString(h.Guesses),
	}
}
