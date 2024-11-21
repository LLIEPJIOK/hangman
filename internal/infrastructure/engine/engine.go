package engine

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"unicode"

	"github.com/LLIEPJIOK/hangman/internal/domain"
)

type Engine struct {
	categories       domain.CategoriesMap
	russianToEnglish map[rune]rune
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
		russianToEnglish: map[rune]rune{
			'А': 'A',
			'В': 'B',
			'Е': 'E',
			'М': 'M',
			'Н': 'H',
			'О': 'O',
			'Р': 'P',
			'С': 'C',
			'Т': 'T',
			'У': 'Y',
			'Х': 'X',
			'а': 'a',
			'е': 'e',
			'о': 'o',
			'р': 'p',
			'с': 'c',
			'у': 'y',
			'х': 'x',
		},
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

func (e *Engine) ToEnglishInLowerCase(letter rune) (rune, bool) {
	if v, ok := e.russianToEnglish[letter]; ok {
		letter = v
	}

	letter = unicode.ToLower(letter)
	if letter < 'a' || letter > 'z' {
		return 0, false
	}

	return letter, true
}

func (e *Engine) CheckLetter(state *domain.GameState, letter rune) {
	var ok bool
	letter, ok = e.ToEnglishInLowerCase(letter)

	// if rune isn't a letter
	if !ok {
		return
	}

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
