package engine

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"unicode"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type Engine struct {
	categories domain.CategoriesMap
}

func New(wordsReader io.Reader) (*Engine, error) {
	var categories map[string]map[string][]domain.Word

	dec := json.NewDecoder(wordsReader)

	err := dec.Decode(&categories)
	if err != nil {
		return nil, fmt.Errorf("cannot parse json to categories map from reader: %w", err)
	}

	categoriesMap := domain.NewCategoriesMap(categories)

	if categoriesMap.Len() == 0 {
		return nil, ErrEmptyWordMap{}
	}

	return &Engine{
		categories: categoriesMap,
	}, nil
}

func (e *Engine) GetRandomWord(category, difficulty string) (domain.Word, error) {
	difficultiesMap, ok := e.categories.Get(category)
	if !ok {
		return domain.Word{}, ErrNoCategory{
			category: category,
		}
	}

	words, ok := difficultiesMap.Get(difficulty)
	if !ok {
		return domain.Word{}, ErrNoDifficulty{
			difficulty: difficulty,
			category:   category,
		}
	}

	maxValue := big.NewInt(int64(len(words)))

	randID, err := rand.Int(rand.Reader, maxValue)
	if err != nil {
		return domain.Word{}, fmt.Errorf("cannot generate random number: %w", err)
	}

	return words[randID.Int64()], nil
}

func (e *Engine) CheckLetter(state *domain.GameState, letter rune) {
	letter = unicode.ToLower(letter)
	runeID, contains, IsWin := 0, false, true

	for _, r := range state.GuessedWord.Value {
		if r == letter {
			state.WordState[runeID] = letter
			contains = true
		} else if state.WordState[runeID] == '-' {
			IsWin = false
		}

		runeID++
	}

	state.IsWin = IsWin
	if !contains {
		state.AttemptsLeft--
	}
}
