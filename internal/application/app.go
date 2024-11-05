package application

import (
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/generators"
	"golang.org/x/term"
)

func Init(cfg *config.Config) {
	width, height, err := term.GetSize(0)
	if err != nil {
		os.Exit(1)
	}

	if cfg.Width < 4 || cfg.Width > width {
		cfg.Width = width / 2
	}

	if cfg.Height < 4 || cfg.Height > height {
		cfg.Height = height / 2
	}

	if cfg.Height < 4 || cfg.Width < 4 {
		os.Exit(1)
	}

	generator := generators.DFSGenerator{}
	maze := generator.Generate(cfg.Width, cfg.Height)
	renderer := domain.DefaultRenderer{}

	fmt.Print(renderer.Render(&maze))
}
