package engine_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/engine"
	"github.com/stretchr/testify/require"
)

const notJSONWords = `
Animals:
  easy:
    - word: "Dog"
      attempts: 6
    - word: "Cat"
      attempts: 6
Fruits:
  medium:
    - word: "Orange"
      attempts: 
`

const noWords = `{
	"Animals": {
		"easy": []
	},
	"Fruits": {}
}
`

func TestNewEngineError(t *testing.T) {
	t.Parallel()

	t.Run("word not in json format", func(t *testing.T) {
		t.Parallel()

		buffer := bytes.NewBuffer([]byte(notJSONWords))
		_, err := engine.New(buffer)
		require.Error(t, err, "engine.New should return error if words not in json format")
	})

	t.Run("empty map", func(t *testing.T) {
		t.Parallel()

		buffer := bytes.NewBuffer([]byte(noWords))
		_, err := engine.New(buffer)
		require.ErrorAs(
			t,
			err,
			&engine.ErrEmptyWordMap{},
			"engine.New should return error for empty map",
		)
	})
}

const words = `{
	"Animals": {
		"easy": [
			{ "word": "Dog", "attempts": 6 },
			{ "word": "Cat", "attempts": 6 }
		],
		"medium": [
			{ "word": "Bear", "attempts": 5 }
		],
		"hard": [
			{ "word": "Dolphin", "attempts": 4 }
		]
	},
	"Fruits": {
		"easy": [
			{ "word": "Apple", "attempts": 6 }
		],
		"medium": [
			{ "word": "Orange", "attempts": 5 }
		],
		"hard": [
			{ "word": "Pineapple", "attempts": 4 }
		]
	},
	"Colors": {
		"easy": [
			{ "word": "Red", "attempts": 5 },
			{ "word": "Blue", "attempts": 6 }
		],
		"hard": [
			{ "word": "Purple", "attempts": 3 }
		]
	},
	"Emotions": {
		"easy": [
			{ "word": "Happy", "attempts": 6 }
		]
	}
}
`

func TestGetRandomWordWithoutError(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	testCases := []struct {
		category      string
		difficulty    string
		suitableWords map[domain.Word]struct{}
	}{
		{
			category:   "animals",
			difficulty: "easy",
			suitableWords: map[domain.Word]struct{}{
				{Value: "dog", Attempts: 6}: {},
				{Value: "cat", Attempts: 6}: {},
			},
		},
		{
			category:   "fruits",
			difficulty: "medium",
			suitableWords: map[domain.Word]struct{}{
				{Value: "orange", Attempts: 5}: {},
			},
		},
		{
			category:   "colors",
			difficulty: "hard",
			suitableWords: map[domain.Word]struct{}{
				{Value: "purple", Attempts: 3}: {},
			},
		},
		{
			category:   "emotions",
			difficulty: "easy",
			suitableWords: map[domain.Word]struct{}{
				{Value: "happy", Attempts: 6}: {},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			t.Parallel()

			word, err := eng.GetRandomWord(testCase.category, testCase.difficulty)
			require.NoError(t, err, "getRandomWord should return without error")
			require.Contains(t, testCase.suitableWords, word, "word should be from correct place")
		})
	}
}

func TestGetRandomWordWithError(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	testCases := []struct {
		category   string
		difficulty string
		err        error
	}{
		{
			category:   "no category",
			difficulty: "easy",
			err:        engine.ErrNoCategory{},
		},
		{
			category:   "animals",
			difficulty: "no difficulty",
			err:        engine.ErrNoDifficulty{},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			t.Parallel()

			_, err := eng.GetRandomWord(testCase.category, testCase.difficulty)
			require.ErrorAs(t, err, &testCase.err, "getRandomWord should return error")
		})
	}
}

func gameStatesEquals(first, second domain.GameState) bool {
	if first.GuessedWord != second.GuessedWord || len(first.WordState) != len(second.WordState) ||
		first.AttemptsLeft != second.AttemptsLeft ||
		first.IsWin != second.IsWin {
		return false
	}

	for i := range len(first.WordState) {
		if first.WordState[i] != second.WordState[i] {
			return false
		}
	}

	return true
}

func TestCheckLetter(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	testCases := []struct {
		initial  domain.GameState
		letter   rune
		expected domain.GameState
	}{
		{
			initial: domain.NewGameState(domain.Word{Value: "word", Attempts: 4}),
			letter:  'w',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "word", Attempts: 4},
				WordState:    []rune("w---"),
				AttemptsLeft: 4,
				IsWin:        false,
			},
		},
		{
			initial: domain.NewGameState(domain.Word{Value: "word", Attempts: 4}),
			letter:  'c',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "word", Attempts: 4},
				WordState:    []rune("----"),
				AttemptsLeft: 3,
				IsWin:        false,
			},
		},
		{
			initial: domain.NewGameState(domain.Word{Value: "aaaaa", Attempts: 4}),
			letter:  'a',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "aaaaa", Attempts: 4},
				WordState:    []rune("aaaaa"),
				AttemptsLeft: 4,
				IsWin:        true,
			},
		},
		{
			initial: domain.GameState{
				GuessedWord:  domain.Word{Value: "word", Attempts: 5},
				WordState:    []rune("w-rd"),
				AttemptsLeft: 1,
				IsWin:        false,
			},
			letter: 'o',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "word", Attempts: 5},
				WordState:    []rune("word"),
				AttemptsLeft: 1,
				IsWin:        true,
			},
		},
		{
			initial: domain.GameState{
				GuessedWord:  domain.Word{Value: "good", Attempts: 10},
				WordState:    []rune("g--d"),
				AttemptsLeft: 10,
				IsWin:        false,
			},
			letter: 'o',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "good", Attempts: 10},
				WordState:    []rune("good"),
				AttemptsLeft: 10,
				IsWin:        true,
			},
		},
		{
			initial: domain.GameState{
				GuessedWord:  domain.Word{Value: "good", Attempts: 10},
				WordState:    []rune("g--d"),
				AttemptsLeft: 1,
				IsWin:        false,
			},
			letter: 'y',
			expected: domain.GameState{
				GuessedWord:  domain.Word{Value: "good", Attempts: 10},
				WordState:    []rune("g--d"),
				AttemptsLeft: 0,
				IsWin:        false,
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			eng.CheckLetter(&testCase.initial, testCase.letter)
			require.Truef(
				t,
				gameStatesEquals(testCase.initial, testCase.expected),
				"game states should be equal: expected: %#v, but got: %#v",
				testCase.initial,
				testCase.expected,
			)
		})
	}
}
