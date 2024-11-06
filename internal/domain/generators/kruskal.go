package generators

import (
	"math/rand"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

const (
	N = iota
	S
	E
	W
)

var (
	DX = map[int]int{E: 1, W: -1, N: 0, S: 0}
	DY = map[int]int{E: 0, W: 0, N: -1, S: 1}
)

type KruskalGenerator struct{}

// Tree structre for union-find algorithm.
type Tree struct {
	Parent *Tree
}

func (t *Tree) Root() *Tree {
	if t.Parent == nil {
		return t
	}

	return t.Parent.Root()
}

func (t *Tree) connected(other *Tree) bool {
	return t.Root() == other.Root()
}

func (t *Tree) connect(other *Tree) {
	other.Root().Parent = t
}

func (g KruskalGenerator) Generate(width, height int) domain.Maze {
	// Create new maze
	maze := domain.NewMaze(width, height)

	// Make tree for each cell
	sets := make([][]*Tree, width)

	for x := 0; x < width; x++ {
		sets[x] = make([]*Tree, height)

		for y := 0; y < height; y++ {
			sets[x][y] = &Tree{}
		}
	}

	var edges [][3]int

	// Boundaries with respect to most top/right/bottom cells to stay unchanged.
	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			if y > 0 && y < height-2 {
				edges = append(edges, [3]int{x, y, S})
			}

			if x > 0 && x < width-2 {
				edges = append(edges, [3]int{x, y, E})
			}
		}
	}

	rand.Shuffle(len(edges), func(i, j int) {
		edges[i], edges[j] = edges[j], edges[i]
	})

	for len(edges) > 0 {
		edge := edges[len(edges)-1]
		edges = edges[:len(edges)-1]

		x, y, direction := edge[0], edge[1], edge[2]
		nx, ny := x+DX[direction], y+DY[direction]

		setA, setB := sets[x][y], sets[nx][ny]

		if !setA.connected(setB) {
			setA.connect(setB)

			switch direction {
			case S:
				maze.Grid[x][y].WallB = 0
			case E:
				maze.Grid[x][y].WallR = 0
			}
		}
	}

	return maze
}
