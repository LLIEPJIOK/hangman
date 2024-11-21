package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/hangman/internal/application/hangman"
)

func main() {
	err := hangman.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot run game: %s", err))
		os.Exit(1)
	}
}
