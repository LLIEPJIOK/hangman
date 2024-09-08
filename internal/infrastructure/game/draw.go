package game

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
)

const (
	gallows = `   ╔═════════╗ 
   ║         ║  
   ║         O  
   ║       / ║ \
   ║        / \  
 ╔╦╩╦╗           
██████████████████`
	maxDrawings = 10
)

func clearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot clear console: %s", err))
	}
}

func (g *Game) draw() {
	clearConsole()

	curDrawings := (g.state.GuessedWord.Attempts - g.state.AttemptsLeft) * maxDrawings / g.state.GuessedWord.Attempts

	for i, line := range g.gallowsLines {
		switch {
		case i+1 == len(g.gallowsLines) || curDrawings-len(g.gallowsLines)-i+1 >= 0:
			fmt.Println(line)
		case curDrawings-len(g.gallowsLines)+i+1 >= 0:
			runes := []rune(line)
			fmt.Println(string(runes[:len(runes)/2]))
		default:
			fmt.Println()
		}
	}

	fmt.Printf("Category: %s, difficulty: %s\n", g.category, g.difficulty)
	fmt.Println("Word:", string(g.state.WordState))
	fmt.Println("Attempts left:", g.state.AttemptsLeft)

	if g.state.IsWin {
		fmt.Println("Congratulation! You guessed the word")
	}

	if g.state.AttemptsLeft == 0 {
		fmt.Println("You lose. Hidden word:", g.state.GuessedWord.Value)
	}
}
