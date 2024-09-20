package game

import (
	"fmt"
)

const greeting = `Welcome to the game "Hangman"!

Here you can test your word knowledge and intuition. Your task is to guess the word letter by letter. 
For each incorrect guess, you'll get closer to failure, so be careful!

You can use the following flags for configuration:
  -c, --category string     words category (default "animals")
  -d, --difficulty string   game difficulty (default "medium")
  -p, --path string         path to file with words (default "resources/words.json")

Good luck and have fun!
`

func (g *Game) drawGreeting() {
	fmt.Fprint(g.out, greeting)
}

const (
	gallows = `   ╔═════════╗ 
   ║         ║  
   ║         O  
   ║       / ║ \
   ║        / \  
 ╔╦╩╦╗           
██████████████████`
	separator   = "\n══════════════════\n\n"
	maxDrawings = 10
)

func (g *Game) draw() {
	curDrawings := (g.state.GuessedWord.Attempts - g.state.AttemptsLeft) * maxDrawings / g.state.GuessedWord.Attempts

	fmt.Fprint(g.out, separator)

	for i, line := range g.gallowsLines {
		switch {
		case i+1 == len(g.gallowsLines) || curDrawings-len(g.gallowsLines)-i+1 >= 0:
			fmt.Fprintln(g.out, line)
		case curDrawings-len(g.gallowsLines)+i+1 >= 0:
			runes := []rune(line)
			fmt.Fprintln(g.out, string(runes[:len(runes)/2]))
		default:
			fmt.Fprintln(g.out)
		}
	}

	fmt.Fprintf(g.out, "Category: %s, difficulty: %s\n", g.category, g.difficulty)
	fmt.Fprintln(g.out, "Word:", string(g.state.WordState))

	if g.showHint {
		fmt.Fprintln(g.out, "Hint:", g.state.GuessedWord.Hint)
	}

	fmt.Fprintln(g.out, "Attempts left:", g.state.AttemptsLeft)

	if g.state.IsWin {
		fmt.Fprintln(g.out, "Congratulations! You guessed the word! You can play again with different parameters")
	}

	if g.state.AttemptsLeft == 0 {
		fmt.Fprintf(g.out, "You lose. Hidden word: %s. Try again!", g.state.GuessedWord.Value)
	}
}
