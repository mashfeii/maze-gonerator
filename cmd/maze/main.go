package main

import (
	"flag"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
)

func main() {
	var width, height int

	var generatorType, solverType string

	var draw bool

	flag.IntVar(&width, "width", 30, "Maze width")
	flag.IntVar(&height, "height", 20, "Maze height")
	flag.StringVar(&generatorType, "generatorType", "dfs", "Algorithm for maze generation (dfs/kruskal)")
	flag.StringVar(&solverType, "solverType", "dfs", "Algorithm to solve the maze (astar/prim)")
	flag.BoolVar(&draw, "draw", false, "Draw the solution")

	flag.Parse()
	application.Init(config.NewConfig(width, height, generatorType, solverType, draw))
}
