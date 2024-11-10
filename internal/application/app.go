package application

import (
	"fmt"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/solvers"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"golang.org/x/term"
)

func validateSize(cfg *config.Config) error {
	width, height, err := term.GetSize(0)
	if err != nil {
		return errors.NewErrTerminalSize()
	}

	if cfg.Width > width {
		cfg.Width = width / 2
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
	renderer := domain.DefaultRenderer{}

	if !ok {
		fmt.Println(errors.NewErrInvalidGenerator(cfg.GeneratorType).Error())
		os.Exit(1)
	}

	maze := generator.Generate(cfg.Width, cfg.Height)

	switch cfg.Command {
	case "draw":
		fmt.Print(renderer.Render(&maze))

	case "solve":
		solver := solvers.AStarSolver{}
		path := solver.Solve(&maze, domain.Coordinate{X: 1, Y: 1}, domain.Coordinate{X: cfg.Width - 2, Y: cfg.Height - 2})
		fmt.Print(renderer.RenderWithPath(&maze, path))

		for _, cell := range path {
			fmt.Printf("%+v\n", maze.Grid[cell.X][cell.Y])
		}

	default:
		fmt.Println(errors.NewErrInvalidCommand(cfg.Command).Error())
		os.Exit(1)
	}
}
