package application

import (
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"golang.org/x/term"
)

func validateSize(cfg *config.Config) error {
	width, height, err := term.GetSize(0)
	if err != nil {
		return errors.NewErrTerminalSize()
	}

	if cfg.Width < 4 || cfg.Width > width {
		cfg.Width = width / 2
	}

	if cfg.Height < 4 || cfg.Height > height {
		cfg.Height = height / 2
	}

	if cfg.Height < 4 || cfg.Width < 4 {
		return errors.NewErrSmallSize(width, height)
	}

	return nil
}

func Init(cfg *config.Config) {
	err := validateSize(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	generatros := config.GetGeneratorTypes()
	generator, ok := generatros[cfg.GeneratorType]

	if !ok {
		fmt.Println(errors.NewErrInvalidGenerator(cfg.GeneratorType).Error())
		os.Exit(1)
	}

	maze := generator.Generate(cfg.Width, cfg.Height)
	renderer := domain.DefaultRenderer{}

	fmt.Print(renderer.Render(&maze))
}
