package solvers

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain"

type Solver interface {
	Solve(maze *domain.Maze, start, end domain.Coordinate) []domain.Coordinate
}
