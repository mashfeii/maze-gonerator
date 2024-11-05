package generators

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain"

type Generator interface {
	Generate(width, height int) domain.Maze
}
