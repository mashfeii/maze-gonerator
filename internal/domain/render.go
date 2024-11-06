package domain

type Renderer interface {
	Render(maze *Maze) string
	RenderWithPath(maze *Maze, path []Coordinate) string
}

type DefaultRenderer struct{}

func (r *DefaultRenderer) Render(maze *Maze) string {
	output := ""
	chars := []rune{' ', '╶', '╷', '╭', '╴', '─', '╮', '┬', '╵', '╰', '│', '├', '╯', '┴', '┤', '┼'}

	for y := 0; y != maze.Height-1; y++ {
		for x := 0; x != maze.Width-1; x++ {
			if x == 0 && y == 0 {
				output += " "
				continue
			}

			if x == maze.Width-2 && y == maze.Height-2 {
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

// func (r *DefaultRenderer) RenderWithPath(maze *Maze, path []Coordinate) string {
// 	output := ""

// 	for rowIdx, row := range maze.Grid {
// 		for colIdx, col := range row {
// 			symbol := ""

// 			if col.Type == Wall {
// 				symbol = "#"
// 			} else if findIndex(path, rowIdx, colIdx) {
// 				symbol = "X"
// 			} else {
// 				symbol = " "
// 			}

// 			output += symbol
// 		}

// 		output += "\n"
// 	}

// 	return output
// }

// func findIndex(arr []Coordinate, x, y int) bool {
// 	for _, v := range arr {
// 		if v.X == x && v.Y == y {
// 			return true
// 		}
// 	}

// 	return false
// }

// func btoi(b bool) int {
// 	if b {
// 		return 1
// 	}

// 	return 0
// }
