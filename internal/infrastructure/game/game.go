package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type Engine interface {
	GetRandomWord(category, difficulty string) (domain.Word, error)
	CheckLetter(state *domain.GameState, letter rune)
}

type Game struct {
	engine       Engine
	state        domain.GameState
	category     string
	difficulty   string
	gallowsLines []string
}

func New(eng Engine, category, difficulty string) (*Game, error) {
	guessedWord, err := eng.GetRandomWord(category, difficulty)
	if err != nil {
		return nil, fmt.Errorf("eng.GetRandomWord(%q, %q): %w", category, difficulty, err)
	}

	return &Game{
		engine:       eng,
		state:        domain.NewGameState(guessedWord),
		category:     category,
		difficulty:   difficulty,
		gallowsLines: strings.Split(gallows, "\n"),
	}, nil
}

func (g *Game) Start() {
	g.draw()

	scan := bufio.NewScanner(os.Stdin)
	for g.state.AttemptsLeft != 0 && !g.state.IsWin {
		if !scan.Scan() {
			break
		}

		input := []rune(scan.Text())
		if len(input) == 1 {
			g.engine.CheckLetter(&g.state, input[0])
		}

		g.draw()
	}
}
