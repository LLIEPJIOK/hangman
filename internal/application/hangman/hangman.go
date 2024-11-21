package hangman

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/LLIEPJIOK/hangman/internal/infrastructure/engine"
	"github.com/LLIEPJIOK/hangman/internal/infrastructure/game"
)

func Run() error {
	var wordsPath, category, difficulty string

	flag.StringVar(&wordsPath, "path", "resources/words.json", "path to file with words")
	flag.StringVar(&wordsPath, "p", "resources/words.json", "path to file with words (shorthand)")
	flag.StringVar(&category, "category", "animals", "words category")
	flag.StringVar(&category, "c", "animals", "words category (shorthand)")
	flag.StringVar(&difficulty, "difficulty", "medium", "game difficulty")
	flag.StringVar(&difficulty, "d", "medium", "game difficulty (shorthand)")

	flag.Parse()

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

	newGame, err := game.New(eng, category, difficulty, os.Stdin, os.Stdout)
	if err != nil {
		return fmt.Errorf("cannot create new game: %w", err)
	}

	newGame.Start()

	return nil
}
