package game

import (
	"bufio"
	"fmt"
	"io"
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
	showHint     bool
	gallowsLines []string
	in           io.Reader
	out          io.Writer
}

func New(eng Engine, category, difficulty string, in io.Reader, out io.Writer) (*Game, error) {
	guessedWord, err := eng.GetRandomWord(category, difficulty)
	if err != nil {
		return nil, fmt.Errorf("eng.GetRandomWord(%q, %q): %w", category, difficulty, err)
	}

	return &Game{
		engine:       eng,
		state:        domain.NewGameState(guessedWord),
		category:     category,
		difficulty:   difficulty,
		showHint:     false,
		gallowsLines: strings.Split(gallows, "\n"),
		in:           in,
		out:          out,
	}, nil
}

func (g *Game) Start() {
	g.drawGreeting()
	g.draw()

	scan := bufio.NewScanner(g.in)
	for g.state.AttemptsLeft != 0 && !g.state.IsWin {
		if !scan.Scan() {
			break
		}

		input := strings.TrimSpace(scan.Text())
		if input == "hint" {
			g.showHint = true
		} else {
			inputRunes := []rune(input)
			if len(input) == 1 {
				g.engine.CheckLetter(&g.state, inputRunes[0])
			}
		}

		g.draw()
	}
}
