package domain

import "strings"

type GameState struct {
	GuessedWord  Word
	WordState    []rune
	AttemptsLeft int
	IsWin        bool
}

func NewGameState(guessedWord Word) GameState {
	return GameState{
		GuessedWord:  guessedWord,
		WordState:    []rune(strings.Repeat("-", len(guessedWord.Value))),
		AttemptsLeft: guessedWord.Attempts,
		IsWin:        false,
	}
}
