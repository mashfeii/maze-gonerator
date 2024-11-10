package generators

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/errors"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

type DFSGenerator struct{}

func (g DFSGenerator) Generate(width, height int) (domain.Maze, error) {
	if width < 4 || height < 4 {
		return domain.Maze{}, errors.NewErrSmallSize(width, height)
	}

	maze := domain.NewMaze(width, height)
	visited := make([][]bool, width)

	for x := 0; x < width; x++ {
		visited[x] = make([]bool, height)
	}

	for y := range height {
		visited[0][y] = true
		visited[width-1][y] = true
	}

	for x := range width {
		visited[x][0] = true
		visited[x][height-1] = true
	}

	visited[0][0] = true
	visited[width-1][0] = true

	startX := pkg.CryptoRandSecure(int64(width))
	startY := pkg.CryptoRandSecure(int64(height))
	newCell := &maze.Grid[startX][startY]

	queue := []*domain.Cell{newCell}
	visited[startX][startY] = true

	for len(queue) > 0 {
		N := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		adjacent := getRandomAdjacent(N, &maze, visited)

		if adjacent != nil {
			breakWall(N, adjacent)

			visited[adjacent.X][adjacent.Y] = true

			queue = append(queue, adjacent, N)
		}
	}

	return maze, nil
}

func breakWall(n, a *domain.Cell) {
	xDiff := n.X - a.X
	yDiff := n.Y - a.Y

	switch {
	case xDiff == 0 && yDiff == -1: // n / a
		n.WallB = 0

	case xDiff == 0 && yDiff == 1: // a / n
		a.WallB = 0

	case xDiff == -1 && yDiff == 0: // n | a
		n.WallR = 0

	case xDiff == 1 && yDiff == 0: // a | n
		a.WallR = 0
	}
}

func BreakWallTesting(n, a *domain.Cell) {
	breakWall(n, a)
}

func checkBoundaries(x, y, height, width int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func getRandomAdjacent(cell *domain.Cell, maze *domain.Maze, visited [][]bool) *domain.Cell {
	adjacent := make([]*domain.Cell, 0)
	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, direction := range directions {
		x := cell.X + direction[0]
		y := cell.Y + direction[1]

		if checkBoundaries(x, y, maze.Height, maze.Width) && !visited[x][y] {
			adjacent = append(adjacent, &maze.Grid[x][y])
		}
	}

	if len(adjacent) > 0 {
		return adjacent[pkg.CryptoRandSecure(int64(len(adjacent)))]
	}

	return nil
}
