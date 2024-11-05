package generators

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain"

type KruskalGenerator struct{}

func (g *KruskalGenerator) Generate(width, height int) domain.Maze {
	return domain.NewMaze(width, height)
}
