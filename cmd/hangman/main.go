package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application/hangman"
)

func main() {
	err := hangman.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("cannot run game: %s", err))
		os.Exit(1)
	}
}
