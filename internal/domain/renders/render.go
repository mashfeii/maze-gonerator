package renders

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/fatih/color"
)

type Renderer interface {
	Render(maze *domain.Maze) string
	RenderWithPath(maze *domain.Maze, path []domain.Coordinate) string
}

type DefaultRenderer struct{}

func (r *DefaultRenderer) Render(maze *domain.Maze) string {
	output := ""
	chars := []rune{' ', '╶', '╷', '╭', '╴', '─', '╮', '┬', '╵', '╰', '│', '├', '╯', '┴', '┤', '┼'}

	for y := 0; y != maze.Height-1; y++ {
		for x := 0; x != maze.Width-1; x++ {
			upper := maze.Grid[x][y].WallR
			left := maze.Grid[x][y].WallB
			down := maze.Grid[x][y+1].WallR
			right := maze.Grid[x+1][y].WallB
			char := upper*8 + left*4 + down*2 + right
			output += string(chars[char])

			if right == 0 {
				output += string(chars[0])
			} else {
				output += string(chars[5])
			}
		}

		output += "\n"
	}

	return output
}

func (r *DefaultRenderer) RenderWithPath(maze *domain.Maze, path []domain.Coordinate) string {
	if len(path) == 0 {
		return r.Render(maze) + color.RedString("No path found")
	}

	output := ""
	chars := []rune{' ', '╶', '╷', '╭', '╴', '─', '╮', '┬', '╵', '╰', '│', '├', '╯', '┴', '┤', '┼'}

	for y := 0; y != maze.Height-1; y++ {
		for x := 0; x != maze.Width-1; x++ {
			if x == path[0].X && y == path[0].Y {
				output += " "
				continue
			}

			if x == path[len(path)-1].X && y == path[len(path)-1].Y {
				output += " "
				continue
			}

			upper := maze.Grid[x][y].WallR
			left := maze.Grid[x][y].WallB
			down := maze.Grid[x][y+1].WallR
			right := maze.Grid[x+1][y].WallB
			char := upper*8 + left*4 + down*2 + right
			output += string(chars[char])

			if right == 0 {
				output += string(chars[0])
			} else {
				output += string(chars[5])
			}
		}

		output += "\n"
	}

	return output
}
