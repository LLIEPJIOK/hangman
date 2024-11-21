package game_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/LLIEPJIOK/hangman/internal/infrastructure/engine"
	"github.com/LLIEPJIOK/hangman/internal/infrastructure/game"
	"github.com/stretchr/testify/require"
)

const words = `{
	"Animals": {
		"easy": [
			{ "word": "Dog", "attempts": 6, "hint": "A loyal pet that barks" }
		]
	},
	"Colors": {
		"hard": [
			{ "word": "Purple", "attempts": 3, "hint": "Mix of blue and red" }
		]
	}
}
`

func TestGameConstructorWithError(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	_, err = game.New(eng, "category", "difficulty", &bytes.Buffer{}, io.Discard)
	require.Error(t, err, "constructor must return error")
}

const (
	inputForWin = `d
a
o
r
hint
g`

	expectedOutputForWin = `Welcome to the game "Hangman"!

Here you can test your word knowledge and intuition. Your task is to guess the word letter by letter. 
For each incorrect guess, you'll get closer to failure, so be careful!

You can use the following flags for configuration:
  -c, --category string     words category (default "animals")
  -d, --difficulty string   game difficulty (default "medium")
  -p, --path string         path to file with words (default "resources/words.json")

You can type "hint" any time you want to get a hint

Good luck and have fun!

══════════════════







██████████████████
Category: animals, difficulty: easy
Word: ---
Attempts left: 6

══════════════════







██████████████████
Category: animals, difficulty: easy
Word: d--
Attempts left: 6

══════════════════






 ╔╦╩╦╗  
██████████████████
Category: animals, difficulty: easy
Word: d--
Attempts left: 5

══════════════════






 ╔╦╩╦╗  
██████████████████
Category: animals, difficulty: easy
Word: do-
Attempts left: 5

══════════════════




   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: animals, difficulty: easy
Word: do-
Attempts left: 4

══════════════════




   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: animals, difficulty: easy
Word: do-
Hint: A loyal pet that barks
Attempts left: 4

══════════════════




   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: animals, difficulty: easy
Word: dog
Hint: A loyal pet that barks
Attempts left: 4
Congratulations! You guessed the word! You can play again with different parameters
`
)

func TestGameWithWin(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	inputBuffer, outputBuffer := bytes.NewBuffer([]byte(inputForWin)), &bytes.Buffer{}
	g, err := game.New(eng, "animals", "easy", inputBuffer, outputBuffer)
	require.NoError(t, err, "game must be created")

	g.Start()

	require.Equal(t, expectedOutputForWin, outputBuffer.String(), "messages should be equal")
}

const (
	inputForLoss = `d
a
hint
p
q`

	expectedOutputForLoss = `Welcome to the game "Hangman"!

Here you can test your word knowledge and intuition. Your task is to guess the word letter by letter. 
For each incorrect guess, you'll get closer to failure, so be careful!

You can use the following flags for configuration:
  -c, --category string     words category (default "animals")
  -d, --difficulty string   game difficulty (default "medium")
  -p, --path string         path to file with words (default "resources/words.json")

You can type "hint" any time you want to get a hint

Good luck and have fun!

══════════════════







██████████████████
Category: colors, difficulty: hard
Word: ------
Attempts left: 3

══════════════════




   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: colors, difficulty: hard
Word: ------
Attempts left: 2

══════════════════

   ╔═════════╗ 
   ║    
   ║    
   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: colors, difficulty: hard
Word: ------
Attempts left: 1

══════════════════

   ╔═════════╗ 
   ║    
   ║    
   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: colors, difficulty: hard
Word: ------
Hint: Mix of blue and red
Attempts left: 1

══════════════════

   ╔═════════╗ 
   ║    
   ║    
   ║    
   ║    
 ╔╦╩╦╗  
██████████████████
Category: colors, difficulty: hard
Word: p--p--
Hint: Mix of blue and red
Attempts left: 1

══════════════════

   ╔═════════╗ 
   ║         ║  
   ║         O  
   ║       / ║ \
   ║        / \  
 ╔╦╩╦╗  
██████████████████
Category: colors, difficulty: hard
Word: p--p--
Hint: Mix of blue and red
Attempts left: 0
You lose. Hidden word: purple. Try again!
`
)

func TestGameWithLoss(t *testing.T) {
	t.Parallel()

	buffer := bytes.NewBuffer([]byte(words))
	eng, err := engine.New(buffer)
	require.NoError(t, err, "engine must be created")

	inputBuffer, outputBuffer := bytes.NewBuffer([]byte(inputForLoss)), &bytes.Buffer{}
	g, err := game.New(eng, "colors", "hard", inputBuffer, outputBuffer)
	require.NoError(t, err, "game must be created")

	g.Start()

	require.Equal(t, expectedOutputForLoss, outputBuffer.String(), "messages should be equal")
}
