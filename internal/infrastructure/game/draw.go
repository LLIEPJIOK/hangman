package game

import (
	"fmt"
)

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
	fmt.Fprintln(g.out, "Attempts left:", g.state.AttemptsLeft)

	if g.state.IsWin {
		fmt.Fprintln(g.out, "Congratulation! You guessed the word")
	}

	if g.state.AttemptsLeft == 0 {
		fmt.Fprintln(g.out, "You lose. Hidden word:", g.state.GuessedWord.Value)
	}
}
