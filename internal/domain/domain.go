package domain

type CellType int

type Cell struct {
	X     int
	Y     int
	WallR int
	WallB int
}

type Maze struct {
	Width  int
	Height int
	Grid   [][]Cell
}

func NewMaze(width, height int) Maze {
	maze := Maze{
		Width:  width,
		Height: height,
	}
	maze.Grid = make([][]Cell, width)

	for x := 0; x < width; x++ {
		maze.Grid[x] = make([]Cell, height)

		for y := 0; y < height; y++ {
			maze.Grid[x][y] = Cell{
				X:     x,
				Y:     y,
				WallR: 1,
				WallB: 1,
			}
		}
	}

	for y := range height {
		maze.Grid[0][y].WallB = 0
		maze.Grid[width-1][y].WallB = 0
	}

	for x := range width {
		maze.Grid[x][0].WallR = 0
		maze.Grid[x][height-1].WallR = 0
	}

	return maze
}

type Coordinate struct {
	X int
	Y int
}

type Solver interface {
	Solve(maze *Maze, start, end Coordinate) []Coordinate
}
