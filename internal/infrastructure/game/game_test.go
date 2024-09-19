package game_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/engine"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/game"
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

	expectedOutputForWin = `
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
Congratulation! You guessed the word
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

	expectedOutputForLoss = `
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
You lose. Hidden word: purple
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
