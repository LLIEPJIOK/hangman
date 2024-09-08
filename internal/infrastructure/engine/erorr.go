package engine

import "fmt"

type ErrEmptyWordMap struct{}

func (e ErrEmptyWordMap) Error() string {
	return "words map is empty"
}

type ErrNoCategory struct {
	category string
}

func (e ErrNoCategory) Error() string {
	return fmt.Sprintf("no category = %q or it's empty", e.category)
}

type ErrNoDifficulty struct {
	difficulty string
	category   string
}

func (e ErrNoDifficulty) Error() string {
	return fmt.Sprintf("no difficulty = %q in category = %q or it's empty", e.difficulty, e.category)
}
