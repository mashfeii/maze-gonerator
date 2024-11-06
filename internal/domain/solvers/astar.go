package solvers

import (
	"container/heap"
	"math"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

type AStarSolver struct{}

type Node struct {
	coord  domain.Coordinate
	gCost  float64
	hCost  float64
	fCost  float64
	parent *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].fCost < pq[j].fCost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}

func (solver AStarSolver) Solve(maze *domain.Maze, start, end domain.Coordinate) []domain.Coordinate {
	openSet := make(map[domain.Coordinate]*Node)
	closedSet := make(map[domain.Coordinate]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)

	startNode := &Node{coord: start, gCost: 0, hCost: heuristic(start, end)}
	startNode.fCost = startNode.gCost + startNode.hCost
	heap.Push(pq, startNode)
	openSet[start] = startNode

	for pq.Len() > 0 {
		currentNode := heap.Pop(pq).(*Node)
		delete(openSet, currentNode.coord)

		closedSet[currentNode.coord] = true

		if currentNode.coord == end {
			return reconstructPath(currentNode)
		}

		for _, neighborCoord := range getNeighbors(currentNode.coord, maze) {
			if closedSet[neighborCoord] {
				continue
			}

			tentativeGCost := currentNode.gCost + 1
			neighborNode, exists := openSet[neighborCoord]

			if !exists || tentativeGCost < neighborNode.gCost {
				if !exists {
					neighborNode = &Node{coord: neighborCoord}
					openSet[neighborCoord] = neighborNode
				}

				neighborNode.gCost = tentativeGCost
				neighborNode.hCost = heuristic(neighborCoord, end)
				neighborNode.fCost = neighborNode.gCost + neighborNode.hCost
				neighborNode.parent = currentNode

				if !exists {
					heap.Push(pq, neighborNode)
				}
			}
		}
	}

	return nil // No path found
}

func heuristic(a, b domain.Coordinate) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

func getNeighbors(coord domain.Coordinate, maze *domain.Maze) []domain.Coordinate {
	x, y := coord.X, coord.Y
	neighbors := []domain.Coordinate{}

	if y < maze.Height-1 && maze.Grid[x][y].WallB == 0 { // Down
		neighbors = append(neighbors, domain.Coordinate{x, y + 1})
	}
	if x < maze.Width-1 && maze.Grid[x][y].WallR == 0 { // Right
		neighbors = append(neighbors, domain.Coordinate{x + 1, y})
	}
	if y > 0 && maze.Grid[x][y-1].WallB == 0 { // Up
		neighbors = append(neighbors, domain.Coordinate{x, y - 1})
	}
	if x > 0 && maze.Grid[x-1][y].WallR == 0 { // Left
		neighbors = append(neighbors, domain.Coordinate{x - 1, y})
	}

	return neighbors
}

func reconstructPath(endNode *Node) []domain.Coordinate {
	var path []domain.Coordinate
	for node := endNode; node != nil; node = node.parent {
		path = append([]domain.Coordinate{node.coord}, path...)
	}

	return path
}
