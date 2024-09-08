package hangman

import (
	"fmt"
	"os"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/engine"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/game"
	"github.com/spf13/pflag"
)

func Run() error {
	var wordsPath string
	var category string
	var difficulty string

	pflag.StringVarP(&wordsPath, "path", "p", "resources/words.json", "path to file with words")
	pflag.StringVarP(&category, "category", "c", "animals", "words category")
	pflag.StringVarP(&difficulty, "difficulty", "d", "medium", "game difficulty")

	pflag.Parse()

	category = strings.ToLower(category)
	difficulty = strings.ToLower(difficulty)

	file, err := os.Open(wordsPath)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", wordsPath, err)
	}

	eng, err := engine.New(file)
	if err != nil {
		return fmt.Errorf("cannot create engine: %w", err)
	}

	game, err := game.New(eng, category, difficulty)
	if err != nil {
		return fmt.Errorf("cannot create game: %w", err)
	}

	game.Start()
	return nil
}
